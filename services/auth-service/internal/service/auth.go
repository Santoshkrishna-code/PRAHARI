package service

import (
	"context"
	"fmt"
	"time"

	"prahari/services/auth-service/internal/domain"
	"prahari/services/auth-service/internal/events"
	"prahari/services/auth-service/internal/repository"
)

type AuthUseCaseImpl struct {
	repo      domain.UserRepository
	cache     repository.UserCache
	publisher events.EventPublisher
	env       string
}

func NewAuthUseCase(
	repo domain.UserRepository,
	cache repository.UserCache,
	publisher events.EventPublisher,
	env string,
) domain.AuthUseCase {
	return &AuthUseCaseImpl{
		repo:      repo,
		cache:     cache,
		publisher: publisher,
		env:       env,
	}
}

func (s *AuthUseCaseImpl) Register(
	ctx context.Context,
	email, password, role, firstName, lastName string,
) (*domain.User, error) {
	// Business validation
	if email == "" || password == "" || role == "" || firstName == "" || lastName == "" {
		return nil, domain.ErrValidationError
	}

	userSub, err := s.repo.SignUp(ctx, email, password, role, firstName, lastName)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:        userSub,
		Email:     email,
		Role:      domain.UserRole(role),
		FirstName: firstName,
		LastName:  lastName,
	}

	// Publish registration event (matching user_registered.json schema)
	topic := fmt.Sprintf("prahari.%s.auth.user-registered.v1", s.env)
	event := map[string]interface{}{
		"event_id":   fmt.Sprintf("evt-%d", time.Now().UnixNano()),
		"timestamp":  time.Now().Format(time.RFC3339),
		"user_id":    user.ID,
		"email":      user.Email,
		"role":       string(user.Role),
		"first_name": user.FirstName,
		"last_name":  user.LastName,
	}
	_ = s.publisher.Publish(ctx, topic, event)

	return user, nil
}

func (s *AuthUseCaseImpl) Login(
	ctx context.Context,
	email, password, ipAddress, userAgent string,
) (*domain.TokenPair, error) {
	if email == "" || password == "" {
		return nil, domain.ErrValidationError
	}

	tokens, err := s.repo.SignIn(ctx, email, password)
	if err != nil {
		return nil, err
	}

	// Publish Login event (matching user_logged_in.json schema)
	topic := fmt.Sprintf("prahari.%s.auth.user-logged-in.v1", s.env)
	event := map[string]interface{}{
		"event_id":   fmt.Sprintf("evt-%d", time.Now().UnixNano()),
		"timestamp":  time.Now().Format(time.RFC3339),
		"user_id":    email, // Initial login user_id context maps to email
		"email":      email,
		"ip_address": ipAddress,
		"user_agent": userAgent,
	}
	_ = s.publisher.Publish(ctx, topic, event)

	return tokens, nil
}

// Stubs to implement domain.AuthUseCase completely
func (s *AuthUseCaseImpl) Refresh(ctx context.Context, refreshToken string) (*domain.TokenPair, error) {
	return s.repo.RefreshToken(ctx, refreshToken)
}

func (s *AuthUseCaseImpl) GetUserProfile(ctx context.Context, accessToken string) (*domain.User, error) {
	return s.repo.GetUserByToken(ctx, accessToken)
}

func (s *AuthUseCaseImpl) VerifyToken(ctx context.Context, token string) (*domain.JWTClaims, error) {
	// Verification logic moves to token verifier helper
	return nil, nil
}
