package handler

import (
	"context"
	"grafana-api/cmd/lambda/common"
)

type LambdaHandler struct {
}

func (h LambdaHandler) HandleRequest(ctx context.Context, req common.ProxyRequest) (resp common.ProxyResponse, err error) {
	return resp, nil
}
