package dto

// RegisterRequest represents the JSON parameters parsed when creating new user profiles.
type RegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	Role      string `json:"role" validate:"required,oneof=Worker SafetySupervisor PlantManager"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}
