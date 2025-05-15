package mapper

import (
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/monitoring/entity"
	"github.com/Ndraaa15/foreglyc-server/pkg/enum"
)

func ToGlucometerMonitoringResponse(data *entity.GlucometerMonitoring) dto.GlucometerMonitoringResponse {
	resp := dto.GlucometerMonitoringResponse{
		Id:           data.Id,
		BloodGlucose: data.BloodGlucose,
		Unit:         data.Unit,
		Status:       data.Status.String(),
		Date:         data.CreatedAt.Time.Format("02 Jan 2006"),
		Time:         data.CreatedAt.Time.Format("15:04"),
		IsSafe:       true,
	}

	if data.Status != enum.GlucoseStatusNormal {
		resp.IsSafe = false
	}

	return resp
}

func ToGlucometerMonitoringGraphResponse(label *string, value *float64, status *string) dto.GlucometerMonitoringGraphResponse {
	return dto.GlucometerMonitoringGraphResponse{
		Label:  *label,
		Value:  *value,
		Status: *status,
	}
}
