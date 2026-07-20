package inspection

import (
	"context"
)

// Repository defines the port contract for inspection storage.
type Repository interface {
	Create(ctx context.Context, inspection *Inspection) error
	FindByID(ctx context.Context, id string) (*Inspection, error)
	FindByNumber(ctx context.Context, number string) (*Inspection, error)
	Update(ctx context.Context, inspection *Inspection) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, offset, limit int) ([]*Inspection, int, error)
}
