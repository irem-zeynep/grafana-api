package domain

import (
	"context"
	"github.com/sirupsen/logrus"
	"grafana-api/domain/model"
	"grafana-api/infrastructure/persistence/timestream"
)

type IAuditService interface {
	SaveAudit(ctx context.Context, dto model.AuditDTO) error
}

type auditService struct {
	repository timestream.IAuditRepository
	logger     *logrus.Logger
}

func NewAuditService(repo timestream.IAuditRepository, logger *logrus.Logger) IAuditService {
	return &auditService{repository: repo, logger: logger}
}

func (s auditService) SaveAudit(ctx context.Context, dto model.AuditDTO) error {
	if err := s.repository.SaveAudit(ctx, dto); err != nil {
		s.logger.Errorf("Could not save audit: %v because: %v", dto, err)
		return err
	}

	return nil
}
