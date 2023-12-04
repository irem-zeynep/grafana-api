package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"grafana-api/cmd/lambda/common"
	"grafana-api/domain"
	"grafana-api/domain/model"
)

type LambdaHandler struct {
	Serv          domain.IGrafanaService
	AuditServ     domain.IAuditService
	ExceptionServ domain.IExceptionService
	Logger        *logrus.Logger
}

func (h *LambdaHandler) HandleRequest(ctx context.Context, req common.ProxyRequest) (proxyResp common.ProxyResponse, err error) {
	defer h.exceptionHandler(ctx, &proxyResp)

	_ = h.AuditServ.SaveAudit(ctx, model.AuditDTO{Method: req.HTTPMethod, Path: req.Path, Payload: req.String()})

	request := model.CheckOrgUserRequest{
		OrgName:   req.QueryStringParameters["org"],
		UserEmail: req.QueryStringParameters["email"],
	}

	resp, err := h.Serv.CheckOrganizationUser(ctx, request)
	if err != nil {
		proxyResp = common.NewProxyErrorResponse(err)
	} else {
		proxyResp = common.ProxyResponse{StatusCode: 200, Body: resp.String()}
	}

	if proxyResp.StatusCode >= 500 {
		_ = h.ExceptionServ.SaveException(ctx, proxyResp.Body)
	}

	return proxyResp, nil
}

func (h *LambdaHandler) exceptionHandler(ctx context.Context, resp *common.ProxyResponse) {
	if e := recover(); e != nil {
		err := errors.New(fmt.Sprint(e))
		_ = h.ExceptionServ.SaveException(ctx, err.Error())
		*resp = common.NewProxyErrorResponse(err)
	}
}
