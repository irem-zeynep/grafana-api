package domain

import (
	"context"
	"github.com/sirupsen/logrus"
	"grafana-api/domain/model"
	"grafana-api/infrastructure/http/grafana"
)

type IGrafanaService interface {
	CheckOrganizationUser(ctx context.Context, cmd model.CheckOrgUserRequest) error
}

type grafanaService struct {
	client grafana.IClient
	logger *logrus.Logger
}

func NewGrafanaService(client grafana.IClient, logger *logrus.Logger) IGrafanaService {
	return &grafanaService{
		client: client,
		logger: logger,
	}
}

func (s grafanaService) CheckOrganizationUser(ctx context.Context, req model.CheckOrgUserRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	s.logger.Infof("Sending request to grafana for organization: %s", req.OrgName)
	_, err := s.client.GetOrg(ctx, req.OrgName)
	if err != nil {
		return err
	}

	s.logger.Infof("Sending request to grafana for user: %s", req.UserEmail)
	_, err = s.client.GetUser(ctx, req.UserEmail)
	if err == nil {
		return nil
	}

	if err.Error() != string(model.NotFoundCode) {
		return err
	}

	s.logger.Infof("Sending request to grafana to create user: %s", req.UserEmail)
	userCmd := grafana.CreateUserCommand{
		Email:    req.UserEmail,
		Name:     req.UserEmail,
		Login:    req.UserEmail,
		Password: "pass",
		Role:     "Viewer",
	}
	if err = s.client.CreateUser(ctx, userCmd); err != nil {
		return err
	}
	s.logger.Infof("Successfully created new user: %s", req.UserEmail)

	return nil
}