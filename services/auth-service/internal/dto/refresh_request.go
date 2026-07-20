package dto

// RefreshRequest wraps the active refresh token payload.
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
