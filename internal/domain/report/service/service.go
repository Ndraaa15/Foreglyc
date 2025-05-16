package service

import (
	"github.com/Ndraaa15/foreglyc-server/internal/domain/report/repository"
	"github.com/sirupsen/logrus"
)

type IReportService interface {
}

type ReportService struct {
	log              *logrus.Logger
	reportRepository repository.RepositoryItf
}

func New(log *logrus.Logger, reportRepository repository.RepositoryItf) IReportService {
	return &ReportService{
		log:              log,
		reportRepository: reportRepository,
	}
}
