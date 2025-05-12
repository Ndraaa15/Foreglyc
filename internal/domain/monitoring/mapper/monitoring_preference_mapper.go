package mapper

import (
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/entity"
)

func ToCGMMonitoringPreferenceResponse(data *entity.CGMMonitoringPreference) dto.CGMMonitoringPrefereceResponse {
	return dto.CGMMonitoringPrefereceResponse{
		UserId:                        data.UserId,
		HypoglycemiaAccuteThreshold:   data.HypoglycemiaAccuteThreshold,
		HypoglycemiaChronicThreshold:  data.HypoglycemiaChronicThreshold,
		HyperglycemiaAccuteThreshold:  data.HyperglycemiaAccuteThreshold,
		HyperglycemiaChronicThreshold: data.HyperglycemiaChronicThreshold,
		SendNotification:              data.SendNotification,
	}
}

func ToGlucometerMonitoringPreferenceResponse(data *entity.GlucometerMonitoringPreference) dto.GlucometerMonitoringPrefereceResponse {
	return dto.GlucometerMonitoringPrefereceResponse{
		UserId:                        data.UserId,
		StartWakeUpTime:               data.StartWakeUpTime.Format("15:04"),
		EndWakeUpTime:                 data.EndWakeUpTime.Format("15:04"),
		PhysicalActivityDays:          data.PhysicalActivityDays,
		StartSleepTime:                data.StartSleepTime.Format("15:04"),
		EndSleepTime:                  data.EndSleepTime.Format("15:04"),
		HypoglycemiaAccuteThreshold:   data.HypoglycemiaAccuteThreshold,
		HypoglycemiaChronicThreshold:  data.HypoglycemiaChronicThreshold,
		HyperglycemiaAccuteThreshold:  data.HyperglycemiaAccuteThreshold,
		HyperglycemiaChronicThreshold: data.HyperglycemiaChronicThreshold,
		SendNotification:              data.SendNotification,
	}
}
