package hazard

import (
	"context"
)

// Repository defines the port contract for hazard storage.
type Repository interface {
	Create(ctx context.Context, hazard *Hazard) error
	FindByID(ctx context.Context, id string) (*Hazard, error)
	FindByNumber(ctx context.Context, number string) (*Hazard, error)
	Update(ctx context.Context, hazard *Hazard) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, offset, limit int) ([]*Hazard, int, error)
}
