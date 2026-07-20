package maintenance

import (
	"context"
)

// Repository defines the port contract for maintenance storage.
type Repository interface {
	Create(ctx context.Context, maintenance *Maintenance) error
	FindByID(ctx context.Context, id string) (*Maintenance, error)
	FindByNumber(ctx context.Context, number string) (*Maintenance, error)
	Update(ctx context.Context, maintenance *Maintenance) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, offset, limit int) ([]*Maintenance, int, error)
}
