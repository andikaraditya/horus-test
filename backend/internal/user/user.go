package user

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username" validate:"required"`
	Name     string `json:"nama"`
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password,omitempty"`
}
