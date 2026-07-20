package compliance

import (
	"context"
)

// Repository defines the contract for compliance storage.
type Repository interface {
	Create(ctx context.Context, compliance *Compliance) error
	FindByID(ctx context.Context, id string) (*Compliance, error)
	FindByNumber(ctx context.Context, number string) (*Compliance, error)
	Update(ctx context.Context, compliance *Compliance) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, offset, limit int) ([]*Compliance, int, error)
}
