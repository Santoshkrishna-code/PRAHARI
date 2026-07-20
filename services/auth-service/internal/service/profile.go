package service

import (
	"context"

	"prahari/services/auth-service/internal/domain"
)

func (s *AuthUseCaseImpl) GetUserProfile(ctx context.Context, accessToken string) (*domain.User, error) {
	if accessToken == "" {
		return nil, domain.ErrValidationError
	}

	// 1. Read-through check: Try to fetch from cache first
	if cachedUser, err := s.cache.Get(ctx, accessToken); err == nil {
		return cachedUser, nil
	}

	// 2. Cache miss: Query Cognito repository directly
	user, err := s.repo.GetUserByToken(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	// 3. Write-through: Populate cache for subsequent requests
	_ = s.cache.Set(ctx, accessToken, user)

	return user, nil
}
