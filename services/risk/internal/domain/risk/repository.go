package risk

import (
	"context"
)

// Repository defines the contract for risk storage.
type Repository interface {
	Create(ctx context.Context, risk *Risk) error
	FindByID(ctx context.Context, id string) (*Risk, error)
	FindByNumber(ctx context.Context, number string) (*Risk, error)
	Update(ctx context.Context, risk *Risk) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, offset, limit int) ([]*Risk, int, error)
}
