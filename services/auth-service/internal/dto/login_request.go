package dto

// LoginRequest defines parameters sent by user during login sessions.
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
