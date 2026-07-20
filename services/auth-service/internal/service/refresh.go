package service

import (
	"context"
	"fmt"
	"time"

	"prahari/services/auth-service/internal/domain"
)

func (s *AuthUseCaseImpl) Refresh(ctx context.Context, refreshToken string) (*domain.TokenPair, error) {
	if refreshToken == "" {
		return nil, domain.ErrValidationError
	}

	tokens, err := s.repo.RefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	// Publish TokenRefreshed event (matching token_refreshed.json schema)
	topic := fmt.Sprintf("prahari.%s.auth.token-refreshed.v1", s.env)
	event := map[string]interface{}{
		"event_id":   fmt.Sprintf("evt-%d", time.Now().UnixNano()),
		"timestamp":  time.Now().Format(time.RFC3339),
		"user_id":    "unknown", // User ID isn't directly retrievable from pure Cognito refresh payload
		"expires_in": tokens.ExpiresIn,
	}
	_ = s.publisher.Publish(ctx, topic, event)

	return tokens, nil
}
