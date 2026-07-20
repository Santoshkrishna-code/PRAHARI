package observation

import (
	"context"
)

// Repository defines the port contract for observation storage.
type Repository interface {
	Create(ctx context.Context, observation *Observation) error
	FindByID(ctx context.Context, id string) (*Observation, error)
	FindByNumber(ctx context.Context, number string) (*Observation, error)
	Update(ctx context.Context, observation *Observation) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, offset, limit int) ([]*Observation, int, error)
}
