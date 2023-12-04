package handler_test

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"grafana-api/cmd/lambda/check-organization-user/handler"
	"grafana-api/cmd/lambda/common"
	"grafana-api/domain/model"
	domainMock "grafana-api/mocks/domain"
	"testing"
)

type HandlerSuite struct {
	suite.Suite
	handler   handler.LambdaHandler
	serv      *domainMock.IGrafanaService
	exServ    *domainMock.IExceptionService
	auditServ *domainMock.IAuditService
	logger    *logrus.Logger
}

func (suite *HandlerSuite) SetupTest() {
	suite.logger = logrus.New()
	suite.serv = domainMock.NewIGrafanaService(suite.T())
	suite.exServ = domainMock.NewIExceptionService(suite.T())
	suite.auditServ = domainMock.NewIAuditService(suite.T())
	suite.handler = handler.LambdaHandler{
		Serv:          suite.serv,
		AuditServ:     suite.auditServ,
		ExceptionServ: suite.exServ,
		Logger:        suite.logger,
	}
}

func (suite *HandlerSuite) TestHandleRequest() {
	//given
	ctx := context.TODO()
	proxyReq := common.ProxyRequest{
		HTTPMethod: "GET",
		Path:       "/test",
		QueryStringParameters: map[string]string{
			"org":   "test org",
			"email": "test email",
		},
	}

	req := model.CheckOrgUserRequest{
		OrgName:   "test org",
		UserEmail: "test email",
	}

	dto := &model.CheckOrgUserDTO{NewUserCreated: false}

	auditDto := model.AuditDTO{Method: "GET", Path: "/test", Payload: proxyReq.String()}
	suite.auditServ.On("SaveAudit", ctx, auditDto).Return(nil)
	suite.serv.On("CheckOrganizationUser", ctx, req).Return(dto, nil)

	//when
	resp, err := suite.handler.HandleRequest(ctx, proxyReq)

	//then
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.StatusCode)
}

func (suite *HandlerSuite) TestHandleRequest_IgnoreErrorWhileSavingAudit() {
	//given
	ctx := context.TODO()
	proxyReq := common.ProxyRequest{
		HTTPMethod: "GET",
		Path:       "/test",
		QueryStringParameters: map[string]string{
			"org":   "test org",
			"email": "test email",
		},
	}

	req := model.CheckOrgUserRequest{
		OrgName:   "test org",
		UserEmail: "test email",
	}
	dto := &model.CheckOrgUserDTO{NewUserCreated: false}

	expectedErr := errors.New("audit err")
	auditDto := model.AuditDTO{Method: "GET", Path: "/test", Payload: proxyReq.String()}
	suite.auditServ.On("SaveAudit", ctx, auditDto).Return(expectedErr)
	suite.serv.On("CheckOrganizationUser", ctx, req).Return(dto, nil)

	//when
	resp, err := suite.handler.HandleRequest(ctx, proxyReq)

	//then
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, resp.StatusCode)
}

func (suite *HandlerSuite) TestHandleRequest_NotFoundError() {
	//given
	ctx := context.TODO()
	proxyReq := common.ProxyRequest{
		HTTPMethod: "GET",
		Path:       "/test",
		QueryStringParameters: map[string]string{
			"org":   "test org",
			"email": "test email",
		},
	}

	req := model.CheckOrgUserRequest{
		OrgName:   "test org",
		UserEmail: "test email",
	}

	appErr := &model.ApplicationError{
		Type: model.NotFoundError,
		Code: string(model.NotFoundCode),
	}
	auditDto := model.AuditDTO{Method: "GET", Path: "/test", Payload: proxyReq.String()}
	suite.auditServ.On("SaveAudit", ctx, auditDto).Return(nil)
	suite.serv.On("CheckOrganizationUser", ctx, req).Return(nil, appErr)

	//when
	resp, err := suite.handler.HandleRequest(ctx, proxyReq)

	//then
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 404, resp.StatusCode)
}

func (suite *HandlerSuite) TestHandleRequest_UnexpectedError() {
	//given
	ctx := context.TODO()
	proxyReq := common.ProxyRequest{
		HTTPMethod: "GET",
		Path:       "/test",
		QueryStringParameters: map[string]string{
			"org":   "test org",
			"email": "test email",
		},
	}

	req := model.CheckOrgUserRequest{
		OrgName:   "test org",
		UserEmail: "test email",
	}

	appErr := &model.ApplicationError{
		Type: model.InternalError,
		Code: string(model.StatusFailedErrorCode),
	}
	auditDto := model.AuditDTO{Method: "GET", Path: "/test", Payload: proxyReq.String()}
	suite.auditServ.On("SaveAudit", ctx, auditDto).Return(nil)
	suite.serv.On("CheckOrganizationUser", ctx, req).Return(nil, appErr)
	suite.exServ.On("SaveException", ctx, "{\"code\":\"client.status-not-success.error\"}").Return(appErr)

	//when
	resp, err := suite.handler.HandleRequest(ctx, proxyReq)

	//then
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 500, resp.StatusCode)
}

func (suite *HandlerSuite) TestHandleRequest_IgnoreErrorWhileSavingException() {
	//given
	ctx := context.TODO()
	proxyReq := common.ProxyRequest{
		HTTPMethod: "GET",
		Path:       "/test",
		QueryStringParameters: map[string]string{
			"org":   "test org",
			"email": "test email",
		},
	}

	req := model.CheckOrgUserRequest{
		OrgName:   "test org",
		UserEmail: "test email",
	}

	appErr := &model.ApplicationError{
		Type: model.InternalError,
		Code: string(model.StatusFailedErrorCode),
	}
	auditDto := model.AuditDTO{Method: "GET", Path: "/test", Payload: proxyReq.String()}

	expectedErr := errors.New("exception saving error")
	suite.auditServ.On("SaveAudit", ctx, auditDto).Return(nil)
	suite.serv.On("CheckOrganizationUser", ctx, req).Return(nil, appErr)
	suite.exServ.On("SaveException", ctx, "{\"code\":\"client.status-not-success.error\"}").Return(expectedErr)

	//when
	resp, err := suite.handler.HandleRequest(ctx, proxyReq)

	//then
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 500, resp.StatusCode)
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}
