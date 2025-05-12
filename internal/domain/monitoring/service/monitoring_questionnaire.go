package service

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/entity"
	"github.com/lib/pq"
	"google.golang.org/genai"
)

func (s *MonitoringService) CreateMonitoringQuestionnaire(ctx context.Context, request dto.CreateMonitoringQuestionnaire, userId string) (dto.MonitoringQuestionnaireResponse, error) {
	repository, err := s.monitoringRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to initialize client")
		return dto.MonitoringQuestionnaireResponse{}, err
	}

	now := time.Now()

	var items []entity.QuestionnaireItem
	for _, item := range request.Questionnaire {
		items = append(items, entity.QuestionnaireItem{
			Question: item.Question,
			Answer:   item.Answer,
		})
	}

	user, err := s.userService.GetUserById(ctx, userId)
	if err != nil {
		s.log.WithError(err).Error("failed to get user")
		return dto.MonitoringQuestionnaireResponse{}, err
	}

	data := &entity.MonitoringQuestionnaire{
		GlucometerMonitoringID: strconv.FormatInt(request.GlucometerMonitoringID, 10),
		Questionnaires:         items,
		ManagementType:         &request.ManagementType,
		CreatedAt:              pq.NullTime{Time: now, Valid: true},
	}

	if err := repository.CreateMonitoringQuestionnaire(ctx, data); err != nil {
		s.log.WithError(err).Error("failed to create monitoring questionnaire")
		return dto.MonitoringQuestionnaireResponse{}, err
	}

	mapData := map[string]any{
		"monitoringQuestionnaire": data,
		"location":                user.Address,
		"caregiverContact":        user.CaregiverContact,
	}

	dataJson, err := json.Marshal(mapData)
	if err != nil {
		s.log.WithError(err).Error("failed to marshal data")
		return dto.MonitoringQuestionnaireResponse{}, err
	}

	message, err := s.geminiAiService.RecomendationAboutQuestionnaire(ctx, genai.Text(string(dataJson)))
	if err != nil {
		return dto.MonitoringQuestionnaireResponse{}, err
	}

	return dto.MonitoringQuestionnaireResponse{
		Message: message,
	}, nil
}
