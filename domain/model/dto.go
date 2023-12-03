package model

type AuditDTO struct {
	Method  string `json:"requestMethod"`
	Path    string `json:"requestPath"`
	Payload string `json:"requestPayload"`
}
