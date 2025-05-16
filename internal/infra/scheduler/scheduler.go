package scheduler

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	foodentity "github.com/Ndraaa15/foreglyc-server/internal/domain/food/entity"
	monitoringentity "github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/entity"
	userentity "github.com/Ndraaa15/foreglyc-server/internal/domain/user/entity"

	"github.com/jmoiron/sqlx"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type IScheduler interface {
	Start()
	InitScheduler()
	Stop()
}

type Scheduler struct {
	db         *sqlx.DB
	cron       *cron.Cron
	log        *logrus.Logger
	n8nWebhook string
}

func New(db *sqlx.DB, log *logrus.Logger, n8nWebhook string) IScheduler {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatalf("failed to load timezone: %v", err)
	}
	return &Scheduler{
		db:         db,
		cron:       cron.New(cron.WithLocation(loc), cron.WithChain(cron.Recover(cron.DefaultLogger))),
		log:        log,
		n8nWebhook: n8nWebhook,
	}
}

func (s *Scheduler) InitScheduler() {
	_, err := s.cron.AddFunc("1 0 * * *", func() {
		tx, err := s.db.BeginTxx(context.Background(), &sql.TxOptions{})
		if err != nil {
			s.log.WithError(err).Error("failed to begin transaction")
			return
		}

		if err := s.createDailyWorkStaff(tx); err != nil {
			tx.Rollback()
			s.log.WithError(err).Error("daily work staff job failed, rolled back")
			return
		}
		if err := tx.Commit(); err != nil {
			s.log.WithError(err).Error("failed to commit transaction")
		}
	})
	if err != nil {
		s.log.WithError(err).Fatal("failed to register cron job")
	}
}

func (s *Scheduler) createDailyWorkStaff(tx *sqlx.Tx) error {
	s.log.Info(">> Starting daily work staff job")

	// 1) load all users
	var users []userentity.User
	if err := tx.Select(&users, `SELECT id, full_name, email, password, photo_profile,
		is_verified, body_weight, date_of_birth, address, caregiver_contact,
		phone_number, auth_provider, level, created_at, updated_at
		FROM users`); err != nil {
		return fmt.Errorf("loading users: %w", err)
	}

	today := time.Now().In(time.FixedZone("Asia/Jakarta", 0)).Format("2006-01-02")

	for _, u := range users {
		userID := u.Id.String()

		// 2) fetch food monitorings
		var foods []foodentity.FoodMonitoring
		if err := tx.Select(&foods, `
			SELECT id, user_id, food_name, meal_time, image_url, nutritions,
			       total_calory, total_carbohydrate, total_protein, total_fat,
					   created_at, updated_at
			FROM food_monitorings
			WHERE user_id = $1 AND DATE(created_at) = $2
		`, userID, today); err != nil {
			s.log.WithError(err).WithField("user", userID).Error("fetching food monitorings")
			continue
		}

		// 3) fetch glucose monitorings
		var glucs []monitoringentity.GlucometerMonitoring
		if err := tx.Select(&glucs, `
			SELECT id, user_id, blood_glucose, unit, status, created_at, updated_at
			FROM glucometer_monitorings
			WHERE user_id = $1 AND DATE(created_at) = $2
		`, userID, today); err != nil {
			s.log.WithError(err).WithField("user", userID).Error("fetching glucometer monitorings")
			continue
		}

		// 4) call n8n
		payload := struct {
			UserID                string                                  `json:"userId"`
			FoodMonitorings       []foodentity.FoodMonitoring             `json:"foodMonitorings"`
			GlucometerMonitorings []monitoringentity.GlucometerMonitoring `json:"glucometerMonitorings"`
		}{
			UserID:                userID,
			FoodMonitorings:       foods,
			GlucometerMonitorings: glucs,
		}
		buf := &bytes.Buffer{}
		if err := json.NewEncoder(buf).Encode(payload); err != nil {
			s.log.WithError(err).WithField("user", userID).Error("encoding webhook payload")
			continue
		}

		resp, err := http.Post(s.n8nWebhook, "application/json", buf)
		if err != nil {
			s.log.WithError(err).WithField("user", userID).Error("posting to n8n webhook")
			continue
		}
		defer resp.Body.Close()

		var result struct {
			Recommendation string `json:"recommendation"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			s.log.WithError(err).WithField("user", userID).Error("decoding webhook response")
			continue
		}

		var totalSugar float64
		for _, g := range glucs {
			totalSugar += g.BloodGlucose
		}

		_, err = tx.Exec(`
			INSERT INTO report_informations
				(user_id, total_blood_sugar, recommendation, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5)
		`, userID, totalSugar, result.Recommendation, time.Now(), time.Now())
		if err != nil {
			s.log.WithError(err).WithField("user", userID).Error("saving report_information")
			continue
		}

		s.log.WithField("user", userID).Info("daily report saved")
	}

	s.log.Info("<< Daily work staff job complete")
	return nil
}

func (s *Scheduler) Start() {
	s.cron.Start()
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
}
