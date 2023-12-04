package model

import (
	"encoding/json"
	"fmt"
)

type AuditDTO struct {
	Method  string `json:"requestMethod"`
	Path    string `json:"requestPath"`
	Payload string `json:"requestPayload"`
}

type CheckOrgUserDTO struct {
	NewUserCreated bool   `json:"newUserCreated"`
	Email          string `json:"email,omitempty"`
	Password       string `json:"password,omitempty"`
}

func (dto CheckOrgUserDTO) String() string {
	rJson, err := json.Marshal(dto)
	if err != nil {
		return fmt.Sprintf("%s %s", dto.Email, dto.Password)
	}

	return string(rJson)

}
