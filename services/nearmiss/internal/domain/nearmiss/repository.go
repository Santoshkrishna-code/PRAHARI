package nearmiss

import (
	"context"
)

// Repository defines the port contract for near miss storage.
type Repository interface {
	Create(ctx context.Context, nearmiss *NearMiss) error
	FindByID(ctx context.Context, id string) (*NearMiss, error)
	FindByNumber(ctx context.Context, number string) (*NearMiss, error)
	Update(ctx context.Context, nearmiss *NearMiss) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, offset, limit int) ([]*NearMiss, int, error)
}
