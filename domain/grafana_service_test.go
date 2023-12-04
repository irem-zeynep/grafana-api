package domain_test

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"grafana-api/domain"
	"grafana-api/domain/model"
	"grafana-api/infrastructure/http/grafana"
	grafanaMock "grafana-api/mocks/infrastructure/http/grafana"
	"testing"
)

type GrafanaServiceSuite struct {
	suite.Suite
	serv   domain.IGrafanaService
	client *grafanaMock.IClient
	logger *logrus.Logger
}

func (suite *GrafanaServiceSuite) SetupTest() {
	suite.logger = logrus.New()
	suite.client = grafanaMock.NewIClient(suite.T())
	suite.serv = domain.NewGrafanaService(suite.client, suite.logger)
}

func (suite *GrafanaServiceSuite) TestSaveGrafana() {
	//given
	ctx := context.TODO()
	req := model.CheckOrgUserRequest{
		OrgName:   "test org name",
		UserEmail: "test user email",
	}

	cmd := grafana.CreateUserCommand{
		Email:    "test user email",
		Name:     "test user email",
		Password: "pass",
		Login:    "test user email",
		Role:     "Viewer",
	}

	orgDto := &grafana.OrganizationDTO{
		ID:   1,
		Name: "orgName",
	}

	suite.client.On("GetOrg", ctx, "test org name").Return(orgDto, nil)
	suite.client.On("GetUser", ctx, "test user email").Return(nil, errors.New(string(model.NotFoundCode)))
	suite.client.On("CreateUser", ctx, cmd).Return(nil)

	//when
	actualResp, err := suite.serv.CheckOrganizationUser(ctx, req)

	//then
	expectedResp := &model.CheckOrgUserDTO{
		NewUserCreated: true,
		Email:          cmd.Email,
		Password:       cmd.Password,
	}

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedResp, actualResp)
}

func (suite *GrafanaServiceSuite) TestSaveGrafana_ErrorWhileCreatingNewUser() {
	//given
	ctx := context.TODO()
	req := model.CheckOrgUserRequest{
		OrgName:   "test org name",
		UserEmail: "test user email",
	}

	cmd := grafana.CreateUserCommand{
		Email:    "test user email",
		Name:     "test user email",
		Password: "pass",
		Login:    "test user email",
		Role:     "Viewer",
	}

	orgDto := &grafana.OrganizationDTO{
		ID:   1,
		Name: "orgName",
	}

	expectedErr := errors.New("expected err")

	suite.client.On("GetOrg", ctx, "test org name").Return(orgDto, nil)
	suite.client.On("GetUser", ctx, "test user email").Return(nil, errors.New(string(model.NotFoundCode)))
	suite.client.On("CreateUser", ctx, cmd).Return(expectedErr)

	//when
	actualResp, err := suite.serv.CheckOrganizationUser(ctx, req)

	//then
	assert.Equal(suite.T(), expectedErr, err)
	assert.Nil(suite.T(), actualResp)
}

func (suite *GrafanaServiceSuite) TestSaveGrafana_UserExists() {
	//given
	ctx := context.TODO()
	req := model.CheckOrgUserRequest{
		OrgName:   "test org name",
		UserEmail: "test user email",
	}

	orgDto := &grafana.OrganizationDTO{
		ID:   1,
		Name: "orgName",
	}

	userDto := &grafana.UserDTO{
		ID:    1,
		Email: "email",
		Login: "login",
	}

	suite.client.On("GetOrg", ctx, "test org name").Return(orgDto, nil)
	suite.client.On("GetUser", ctx, "test user email").Return(userDto, nil)

	//when
	actualResp, err := suite.serv.CheckOrganizationUser(ctx, req)

	//then
	expectedResp := &model.CheckOrgUserDTO{
		NewUserCreated: false,
	}

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedResp, actualResp)
}

func (suite *GrafanaServiceSuite) TestSaveGrafana_ErrorWhileGettingUser() {
	//given
	ctx := context.TODO()
	req := model.CheckOrgUserRequest{
		OrgName:   "test org name",
		UserEmail: "test user email",
	}

	orgDto := &grafana.OrganizationDTO{
		ID:   1,
		Name: "orgName",
	}

	expectedErr := errors.New("expected err")

	suite.client.On("GetOrg", ctx, "test org name").Return(orgDto, nil)
	suite.client.On("GetUser", ctx, "test user email").Return(nil, expectedErr)

	//when
	actualResp, err := suite.serv.CheckOrganizationUser(ctx, req)

	//then
	assert.Equal(suite.T(), expectedErr, err)
	assert.Nil(suite.T(), actualResp)
}

func (suite *GrafanaServiceSuite) TestSaveGrafana_ErrorWhileGettingOrg() {
	//given
	ctx := context.TODO()
	req := model.CheckOrgUserRequest{
		OrgName:   "test org name",
		UserEmail: "test user email",
	}

	expectedErr := errors.New("expected err")

	suite.client.On("GetOrg", ctx, "test org name").Return(nil, expectedErr)

	//when
	actualResp, err := suite.serv.CheckOrganizationUser(ctx, req)

	//then
	assert.Equal(suite.T(), expectedErr, err)
	assert.Nil(suite.T(), actualResp)
}

func (suite *GrafanaServiceSuite) TestSaveGrafana_UserEmailEmpty() {
	//given
	ctx := context.TODO()
	req := model.CheckOrgUserRequest{
		OrgName: "test org name",
	}

	//when
	actualResp, err := suite.serv.CheckOrganizationUser(ctx, req)

	//then
	assert.Equal(suite.T(), string(model.MissingEmailParam), err.Error())
	assert.Nil(suite.T(), actualResp)
}

func (suite *GrafanaServiceSuite) TestSaveGrafana_OrgNameEmpty() {
	//given
	ctx := context.TODO()
	req := model.CheckOrgUserRequest{
		UserEmail: "test user email",
	}

	//when
	actualResp, err := suite.serv.CheckOrganizationUser(ctx, req)

	//then
	assert.Equal(suite.T(), string(model.MissingOrgParam), err.Error())
	assert.Nil(suite.T(), actualResp)
}

func TestGrafanaServiceSuite(t *testing.T) {
	suite.Run(t, new(GrafanaServiceSuite))
}
