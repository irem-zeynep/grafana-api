package grafana

type CreateUserCommand struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Login    string `json:"login"`
	Role     string `json:"role"`
}
