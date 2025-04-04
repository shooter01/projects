package auth

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	LoginResponse
}

type RegisterRequest struct {
	Name string `json:"name" validate:"required"`
	LoginRequest
}
