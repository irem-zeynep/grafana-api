package handler

import (
	"context"
	"github.com/sirupsen/logrus"
	"grafana-api/cmd/lambda/common"
	"grafana-api/domain"
	"grafana-api/domain/model"
)

type LambdaHandler struct {
	Serv   domain.IGrafanaService
	Logger *logrus.Logger
}

func (h LambdaHandler) HandleRequest(ctx context.Context, req common.ProxyRequest) (resp common.ProxyResponse, err error) {
	request := model.CheckOrgUserRequest{
		OrgName:   req.QueryStringParameters["org"],
		UserEmail: req.QueryStringParameters["email"],
	}

	if err = h.Serv.CheckOrganizationUser(ctx, request); err != nil {
		resp = common.NewProxyErrorResponse(err)
	} else {
		resp = common.ProxyResponse{StatusCode: 200}
	}

	return resp, nil
}
