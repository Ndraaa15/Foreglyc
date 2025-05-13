package service

import (
	"context"
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/entity"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/mapper"
	"github.com/Ndraaa15/foreglyc-server/pkg/constant"
	"github.com/lib/pq"
)

func (s *MonitoringService) CreateGlucometerMonitoring(ctx context.Context, request dto.CreateGlucometerMonitoringRequest, userId string) (dto.GlucometerMonitoringResponse, error) {
	repository, err := s.monitoringRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to initialize client")
		return dto.GlucometerMonitoringResponse{}, err
	}

	data := entity.GlucometerMonitoring{
		UserId:       userId,
		BloodGlucose: request.BloodGlucose,
		Status:       BloodGlucoseStatus(request.BloodGlucose),
		Unit:         constant.MG_DL,
		CreatedAt:    pq.NullTime{Time: time.Now(), Valid: true},
	}

	err = repository.CreateGlucometerMonitoring(ctx, &data)
	if err != nil {
		s.log.WithError(err).Error("failed to create glucometer monitoring")
		return dto.GlucometerMonitoringResponse{}, err
	}

	return mapper.ToGlucometerMonitoringResponse(&data), nil
}

func (s *MonitoringService) GetGlucometerMonitorings(ctx context.Context, filter dto.GetGlucometerMonitoringFilter) ([]dto.GlucometerMonitoringResponse, error) {
	repository, err := s.monitoringRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to create transaction")
		return []dto.GlucometerMonitoringResponse{}, err
	}

	data, err := repository.GetGlucometerMonitorings(ctx, filter)
	if err != nil {
		s.log.WithError(err).Error("failed to get glucometer monitorings")
		return []dto.GlucometerMonitoringResponse{}, err
	}

	var resp []dto.GlucometerMonitoringResponse
	for _, item := range data {
		resp = append(resp, mapper.ToGlucometerMonitoringResponse(&item))
	}

	return resp, nil
}

func (s *MonitoringService) GetGlucometerMonitorignGraph(ctx context.Context, filter dto.GetGlucometerMonitoringGraphFilter) ([]dto.GlucometerMonitoringGraphResponse, error) {
	repository, err := s.monitoringRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to create transaction")
		return []dto.GlucometerMonitoringGraphResponse{}, err
	}

	var glucometerMonitoringFilter dto.GetGlucometerMonitoringFilter
	if filter.Type == constant.GlucoseMonitoringHourly {
		glucometerMonitoringFilter.UserId = filter.UserId
		glucometerMonitoringFilter.StartDate = time.Now()
		glucometerMonitoringFilter.EndDate = time.Now()
	} else if filter.Type == constant.GlucoseMonitoringDaily {
		glucometerMonitoringFilter.UserId = filter.UserId
		glucometerMonitoringFilter.StartDate = time.Now()
		glucometerMonitoringFilter.EndDate = time.Now().AddDate(0, 0, -7)
	}

	data, err := repository.GetGlucometerMonitorings(ctx, glucometerMonitoringFilter)
	if err != nil {
		s.log.WithError(err).Error("failed to get glucometer monitorings")
		return []dto.GlucometerMonitoringGraphResponse{}, err
	}

	mapGraph := make(map[string]float64)
	if filter.Type == constant.GlucoseMonitoringHourly {
		for _, item := range data {
			time := item.CreatedAt.Time.Format("15:04")
			if _, ok := mapGraph[time]; !ok {
				mapGraph[time] = 0
			}
			mapGraph[time] += item.BloodGlucose
		}
	} else if filter.Type == constant.GlucoseMonitoringDaily {
		for _, item := range data {
			date := item.CreatedAt.Time.Format("02 Jan 2006")
			if _, ok := mapGraph[date]; !ok {
				mapGraph[date] = 0
			}
			mapGraph[date] += item.BloodGlucose
		}
	}

	var resp []dto.GlucometerMonitoringGraphResponse
	for key, value := range mapGraph {
		resp = append(resp, mapper.ToGlucometerMonitoringGraphResponse(&key, &value))
	}

	return resp, nil
}

func (s *MonitoringService) GetGlucometerMonitoringIds(ctx context.Context, userId string) ([]int64, error) {
	repository, err := s.monitoringRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to create transaction")
		return []int64{}, err
	}

	return repository.GetGlucometerMonitoringIds(ctx, userId)
}
