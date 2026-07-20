package audit

import (
	"context"
)

// Repository defines the contract for audit storage.
type Repository interface {
	Create(ctx context.Context, audit *Audit) error
	FindByID(ctx context.Context, id string) (*Audit, error)
	FindByNumber(ctx context.Context, number string) (*Audit, error)
	Update(ctx context.Context, audit *Audit) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, offset, limit int) ([]*Audit, int, error)
}
