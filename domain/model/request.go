package model

import (
	"time"
)

type CheckOrgUserRequest struct {
	OrgName   string `json:"orgName"`
	UserEmail string `json:"userEmail"`
}

func (r CheckOrgUserRequest) Validate() error {
	if r.OrgName == "" {
		return &ApplicationError{
			Instant: time.Now().String(),
			Code:    string(MissingOrgParam),
			Type:    ValidationError,
		}
	}

	if r.UserEmail == "" {
		return &ApplicationError{
			Instant: time.Now().String(),
			Code:    string(MissingEmailParam),
			Type:    ValidationError,
		}
	}

	return nil
}
