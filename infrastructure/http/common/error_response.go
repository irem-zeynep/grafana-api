package common

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"grafana-api/domain/model"
	"time"
)

func NewErrorResponse(resp *fasthttp.Response, errorOwnerID string) *model.ApplicationError {
	errorType := model.InternalError
	errCode := string(model.StatusFailedErrorCode)
	if resp.StatusCode() == 404 {
		errorType = model.NotFoundError
		errCode = fmt.Sprintf(string(model.NotFoundCode), errorOwnerID)
	}

	return &model.ApplicationError{
		Type:    errorType,
		Instant: time.Now().String(),
		Code:    errCode,
		Cause:   fmt.Sprintf("%s Http client status is not success. Response:\n %s", errorOwnerID, resp.String()),
	}
}

func IsFailure(resp *fasthttp.Response) bool {
	return resp.StatusCode() > 399
}
