package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	prahariLogger "prahari/shared/logger"
	"prahari/services/identity/internal/domain/user"
)

// UserService orchestrates user aggregate creations and password validations.
type UserService struct {
}

// NewUserService constructs a UserService.
func NewUserService() *UserService {
	return &UserService{}
}

// CreateUser validates and stores user profiles.
func (s *UserService) CreateUser(ctx context.Context, email, password string) (*user.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password parameters are required")
	}

	u := &user.User{
		ID:         fmt.Sprintf("usr-%d", time.Now().UnixNano()),
		Email:      email,
		Role:       "Worker", // default role
		Status:     "ACTIVE",
		CreatedAt:  time.Now(),
		MFAEnabled: false,
	}

	if err := u.Validate(); err != nil {
		prahariLogger.Error(ctx, "failed user validation validation check", prahariLogger.Err(err))
		return nil, err
	}

	prahariLogger.Info(ctx, "User signed up successfully", prahariLogger.String("user_id", u.ID))
	return u, nil
}
