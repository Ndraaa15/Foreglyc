package mapper

import (
	"github.com/Ndraaa15/foreglyc-server/internal/domain/report/dto"
	"github.com/Ndraaa15/foreglyc-server/internal/domain/report/entity"
)

func ReportInformationToResponse(data *entity.ReportInformation) dto.ReportInformationResponse {
	return dto.ReportInformationResponse{
		Id:                         data.Id,
		Date:                       data.CreatedAt.Time.Format("02 Jan 2006"),
		TotalBloodGlucose:          data.TotalBloodGlucose,
		RecommendationBloodGlucose: data.RecommendationBloodGlucose,
		Recommendation:             data.Recommendation,
	}
}
