package permit

import (
	"context"
)

// Repository defines the port contract for permit storage.
type Repository interface {
	Create(ctx context.Context, permit *Permit) error
	FindByID(ctx context.Context, id string) (*Permit, error)
	FindByNumber(ctx context.Context, number string) (*Permit, error)
	Update(ctx context.Context, permit *Permit) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, offset, limit int) ([]*Permit, int, error)
}
