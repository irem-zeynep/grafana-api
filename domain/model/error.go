package model

type ErrorCode string

const (
	NotFoundCode          ErrorCode = "domain.not-found.error"
	StatusFailedErrorCode ErrorCode = "client.status-not-success.error"
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
	Code    ErrorCode `json:"code,omitempty"`
	Cause   string    `json:"-"`
}

func (e ApplicationError) Error() string {
	return string(e.Code)
}
