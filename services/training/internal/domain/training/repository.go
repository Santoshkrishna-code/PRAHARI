package training

import (
	"context"
)

// Repository defines the contract for training storage.
type Repository interface {
	Create(ctx context.Context, training *Training) error
	FindByID(ctx context.Context, id string) (*Training, error)
	FindByNumber(ctx context.Context, number string) (*Training, error)
	Update(ctx context.Context, training *Training) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, offset, limit int) ([]*Training, int, error)
}
