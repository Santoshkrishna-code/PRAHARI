package asset

import (
	"context"
)

// Repository defines the port contract for asset storage.
type Repository interface {
	Create(ctx context.Context, asset *Asset) error
	FindByID(ctx context.Context, id string) (*Asset, error)
	FindByNumber(ctx context.Context, number string) (*Asset, error)
	Update(ctx context.Context, asset *Asset) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, offset, limit int) ([]*Asset, int, error)
}
