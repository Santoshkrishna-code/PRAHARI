package contractor

import (
	"context"
)

// Repository defines the port contract for contractor storage.
type Repository interface {
	Create(ctx context.Context, contractor *Contractor) error
	FindByID(ctx context.Context, id string) (*Contractor, error)
	FindByNumber(ctx context.Context, number string) (*Contractor, error)
	Update(ctx context.Context, contractor *Contractor) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, offset, limit int) ([]*Contractor, int, error)
}
