package dto

type RegisterRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}
