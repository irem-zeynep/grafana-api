package common

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"grafana-api/domain/model"
	"time"
)

func NewErrorResponse(resp *fasthttp.Response) *model.ApplicationError {
	errorType := model.InternalError
	errCode := model.StatusFailedErrorCode
	if resp.StatusCode() == 404 {
		errorType = model.NotFoundError
		errCode = model.NotFoundCode
	}

	return &model.ApplicationError{
		Type:    errorType,
		Instant: time.Now().String(),
		Code:    errCode,
		Cause:   fmt.Sprintf("Http client status is not success. grafana:\n %s", resp.String()),
	}
}

func IsFailure(resp *fasthttp.Response) bool {
	return resp.StatusCode() > 399
}
