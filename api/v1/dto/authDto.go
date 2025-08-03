package dto

type RegisterRequest struct {
	Email string `json:"email" required:"true"`
}

type ResetPasswordRequest struct {
	Password string `json:"password" required:"true"`
}

type LoginRequest struct {
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}

type LoginResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}
