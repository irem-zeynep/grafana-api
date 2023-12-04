package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"grafana-api/domain/model"
	"time"
)

type ProxyRequest events.APIGatewayProxyRequest
type ProxyResponse events.APIGatewayProxyResponse

type ErrorBody struct {
	Instant string `json:"instant,omitempty"`
	Code    string `json:"code,omitempty"`
	Cause   string `json:"cause,omitempty"`
}

func (e ErrorBody) Error() string {
	return e.Code
}

func (e ErrorBody) String() string {
	jsonBytes, err := json.Marshal(e)
	if err != nil {
		return e.Code
	}

	return string(jsonBytes)
}

func NewProxyErrorResponse(err error) ProxyResponse {
	var appErr *model.ApplicationError
	if errors.As(err, &appErr) {
		return ProxyResponse{
			StatusCode: getStatusCode(appErr),
			Body: ErrorBody{
				Instant: appErr.Instant,
				Code:    appErr.Code,
				Cause:   appErr.Cause,
			}.String(),
		}
	}

	return ProxyResponse{
		StatusCode: 500,
		Body: ErrorBody{
			Instant: time.Now().String(),
			Code:    "unexpected.system.error",
			Cause:   err.Error(),
		}.String(),
	}
}

func getStatusCode(response *model.ApplicationError) int {
	switch errorType := response.Type; errorType {
	case model.ValidationError:
		return 400
	case model.NotFoundError:
		return 404
	default:
		return 500
	}
}

func (r ProxyRequest) String() string {
	rJson, err := json.Marshal(r)
	if err != nil {
		return fmt.Sprintf("%s %s %s", r.HTTPMethod, r.Path, r.Body)
	}

	return string(rJson)
}
