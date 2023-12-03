package grafana

type OrganizationDTO struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type UserDTO struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	Login string `json:"login"`
}
