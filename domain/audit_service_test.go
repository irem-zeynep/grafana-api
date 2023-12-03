package domain_test

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"grafana-api/domain"
	"grafana-api/domain/model"
	timestreamMock "grafana-api/mocks/infrastructure/persistence/timestream"
	"testing"
)

type AuditServiceSuite struct {
	suite.Suite
	serv   domain.IAuditService
	repo   *timestreamMock.IAuditRepository
	logger *logrus.Logger
}

func (suite *AuditServiceSuite) SetupTest() {
	suite.logger = logrus.New()
	suite.repo = timestreamMock.NewIAuditRepository(suite.T())
	suite.serv = domain.NewAuditService(suite.repo, suite.logger)
}

func (suite *AuditServiceSuite) TestSaveAudit() {
	ctx := context.TODO()
	dto := model.AuditDTO{
		Method:  "GET",
		Path:    "/organization-users",
		Payload: "test payload",
	}
	suite.repo.On("SaveAudit", ctx, dto).Return(nil)
	err := suite.serv.SaveAudit(ctx, dto)

	assert.Nil(suite.T(), err)
}

func (suite *AuditServiceSuite) TestSaveAudit_ErrorWhileSaving() {
	ctx := context.TODO()
	dto := model.AuditDTO{
		Method:  "GET",
		Path:    "/organization-users",
		Payload: "test payload",
	}
	expectedErr := errors.New("expected err")
	suite.repo.On("SaveAudit", ctx, dto).Return(expectedErr)
	err := suite.serv.SaveAudit(ctx, dto)

	assert.Equal(suite.T(), expectedErr, err)
}

func TestAuditServiceSuite(t *testing.T) {
	suite.Run(t, new(AuditServiceSuite))
}
