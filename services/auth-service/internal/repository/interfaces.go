package repository

import (
	"context"

	"prahari/services/auth-service/internal/domain"
)

// UserCache defines outbound caching port contract.
type UserCache interface {
	Get(ctx context.Context, tokenStr string) (*domain.User, error)
	Set(ctx context.Context, tokenStr string, user *domain.User) error
	Delete(ctx context.Context, tokenStr string) error
}
