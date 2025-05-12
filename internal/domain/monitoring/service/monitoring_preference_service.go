package service

import (
	"context"
	"time"

	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/entity"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/mapper"
	"github.com/lib/pq"
)

func (s *MonitoringService) CreateCGMMonitoringPreference(ctx context.Context, request dto.CreateCGMMonitoringPreferenceRequest, userId string) (dto.CGMMonitoringPrefereceResponse, error) {
	repository, err := s.monitoringRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to create transaction")
		return dto.CGMMonitoringPrefereceResponse{}, err
	}

	data := entity.CGMMonitoringPreference{
		UserId:                        userId,
		HypoglycemiaAccuteThreshold:   request.HypoglycemiaAccuteThreshold,
		HypoglycemiaChronicThreshold:  request.HypoglycemiaChronicThreshold,
		HyperglycemiaAccuteThreshold:  request.HyperglycemiaAccuteThreshold,
		HyperglycemiaChronicThreshold: request.HyperglycemiaChronicThreshold,
		SendNotification:              *request.SendNotification,
		CreatedAt:                     pq.NullTime{Time: time.Now(), Valid: true},
	}

	err = repository.CreateCGMMonitoringPreference(ctx, &data)
	if err != nil {
		s.log.WithError(err).Error("failed to create CGM monitoring preference")
		return dto.CGMMonitoringPrefereceResponse{}, err
	}

	return mapper.ToCGMMonitoringPreferenceResponse(&data), nil
}

func (s *MonitoringService) GetCGMMonitoringPreference(ctx context.Context, userId string) (dto.CGMMonitoringPrefereceResponse, error) {
	repository, err := s.monitoringRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to create transaction")
		return dto.CGMMonitoringPrefereceResponse{}, err
	}

	cgmMonitoring, err := repository.GetCGMMonitoringPreference(ctx, userId)
	if err != nil {
		s.log.WithError(err).Error("failed to get CGM monitoring preference")
		return dto.CGMMonitoringPrefereceResponse{}, err
	}

	return mapper.ToCGMMonitoringPreferenceResponse(&cgmMonitoring), nil
}

func (s *MonitoringService) CreateGlucometerMonitoringPreference(ctx context.Context, request dto.CreateGlucometerMonitoringPreferenceRequest, userId string) (dto.GlucometerMonitoringPrefereceResponse, error) {
	repository, err := s.monitoringRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to create transaction")
		return dto.GlucometerMonitoringPrefereceResponse{}, err
	}

	startWakeUpTime, err := time.Parse("15:04", request.StartWakeUpTime)
	if err != nil {
		s.log.WithError(err).Error("failed to parse start wake up time")
		return dto.GlucometerMonitoringPrefereceResponse{}, err
	}

	endWakeUpTime, err := time.Parse("15:04", request.EndWakeUpTime)
	if err != nil {
		s.log.WithError(err).Error("failed to parse end wake up time")
		return dto.GlucometerMonitoringPrefereceResponse{}, err
	}

	startSleepTime, err := time.Parse("15:04", request.StartSleepTime)
	if err != nil {
		s.log.WithError(err).Error("failed to parse start sleep time")
		return dto.GlucometerMonitoringPrefereceResponse{}, err
	}

	endSleepTime, err := time.Parse("15:04", request.EndSleepTime)
	if err != nil {
		s.log.WithError(err).Error("failed to parse end sleep time")
		return dto.GlucometerMonitoringPrefereceResponse{}, err
	}

	data := entity.GlucometerMonitoringPreference{
		UserId:                        userId,
		StartWakeUpTime:               startWakeUpTime,
		EndWakeUpTime:                 endWakeUpTime,
		PhysicalActivityDays:          request.PhysicalActivityDays,
		StartSleepTime:                startSleepTime,
		EndSleepTime:                  endSleepTime,
		HypoglycemiaAccuteThreshold:   request.HypoglycemiaAccuteThreshold,
		HypoglycemiaChronicThreshold:  request.HypoglycemiaChronicThreshold,
		HyperglycemiaAccuteThreshold:  request.HyperglycemiaAccuteThreshold,
		HyperglycemiaChronicThreshold: request.HyperglycemiaChronicThreshold,
		SendNotification:              *request.SendNotification,
		CreatedAt:                     pq.NullTime{Time: time.Now(), Valid: true},
	}

	err = repository.CreateGlucometerMonitoringPreference(ctx, &data)
	if err != nil {
		s.log.WithError(err).Error("failed to create glucometer monitoring preference")
		return dto.GlucometerMonitoringPrefereceResponse{}, err
	}

	return mapper.ToGlucometerMonitoringPreferenceResponse(&data), nil
}

func (s *MonitoringService) GetGlucometerMonitoringPreference(ctx context.Context, userId string) (dto.GlucometerMonitoringPrefereceResponse, error) {
	repository, err := s.monitoringRepository.WithTx(false)
	if err != nil {
		s.log.WithError(err).Error("failed to create transaction")
		return dto.GlucometerMonitoringPrefereceResponse{}, err
	}

	glucometerMonitoring, err := repository.GetGlucometerMonitoringPreference(ctx, userId)
	if err != nil {
		s.log.WithError(err).Error("failed to get glucometer monitoring preference")
		return dto.GlucometerMonitoringPrefereceResponse{}, err
	}

	return mapper.ToGlucometerMonitoringPreferenceResponse(&glucometerMonitoring), nil
}
