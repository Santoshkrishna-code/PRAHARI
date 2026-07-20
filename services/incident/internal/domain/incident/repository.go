package incident

import (
	"context"
)

// Repository defines the port interface for incident persistence operations.
// Infrastructure adapters (PostgreSQL, in-memory) implement this contract.
type Repository interface {
	// Create persists a new incident aggregate.
	Create(ctx context.Context, incident *Incident) error

	// FindByID retrieves an incident by its unique identifier.
	FindByID(ctx context.Context, id string) (*Incident, error)

	// FindByNumber retrieves an incident by its human-readable incident number.
	FindByNumber(ctx context.Context, number string) (*Incident, error)

	// Update persists modifications to an existing incident.
	Update(ctx context.Context, incident *Incident) error

	// Delete performs a soft delete on an incident.
	Delete(ctx context.Context, id string) error

	// List retrieves a paginated list of incidents with optional filtering.
	List(ctx context.Context, offset, limit int) ([]*Incident, int, error)
}
