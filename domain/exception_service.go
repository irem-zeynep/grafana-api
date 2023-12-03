package domain

import (
	"context"
	"github.com/sirupsen/logrus"
	"grafana-api/infrastructure/event/sns"
)

type IExceptionService interface {
	SaveException(ctx context.Context, exceptionMsg string) error
}

type exceptionService struct {
	publisher sns.IEventPublisher
	logger    *logrus.Logger
}

func NewExceptionService(publisher sns.IEventPublisher, logger *logrus.Logger) IExceptionService {
	return &exceptionService{publisher: publisher, logger: logger}
}

func (s exceptionService) SaveException(ctx context.Context, exceptionMsg string) error {
	s.logger.Info("Going to send error message")
	if err := s.publisher.SendMessage(ctx, exceptionMsg); err != nil {
		s.logger.Errorf("Could not send error message: %s because: %v", exceptionMsg, err)
		return err
	}

	return nil
}
