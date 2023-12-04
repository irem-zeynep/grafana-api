package model

type ErrorCode string

const (
	NotFoundCode          ErrorCode = "not-found.%s.error"
	StatusFailedErrorCode ErrorCode = "client.status-not-success.error"
	MissingOrgParam       ErrorCode = "validation.missing-required-param-org.error"
	MissingEmailParam     ErrorCode = "validation.missing-required-param-email.error"
)

type ErrorType string

const (
	NotFoundError   ErrorType = "NOT_FOUND"
	ValidationError ErrorType = "VALIDATION"
	InternalError   ErrorType = "INTERNAL"
)

type ApplicationError struct {
	Instant string    `json:"instant,omitempty"`
	Type    ErrorType `json:"statusCode,omitempty"`
	Code    string    `json:"code,omitempty"`
	Cause   string    `json:"-"`
}

func (e *ApplicationError) Error() string {
	return string(e.Code)
}
