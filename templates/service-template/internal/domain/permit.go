package domain

import (
	"context"
	"time"
)

// Permit represents the Permit to Work domain entity.
type Permit struct {
	ID         string    `json:"id" db:"id"`
	WorkerID   string    `json:"worker_id" db:"worker_id"`
	ZoneID     string    `json:"zone_id" db:"zone_id"`
	Status     string    `json:"status" db:"status"` // PENDING, APPROVED, EXPIRED, REVOKED
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	ExpiresAt  time.Time `json:"expires_at" db:"expires_at"`
	ApprovedBy string    `json:"approved_by" db:"approved_by"`
}

// PermitRepository defines the outbound port for data storage.
type PermitRepository interface {
	GetByID(ctx context.Context, id string) (*Permit, error)
	Create(ctx context.Context, permit *Permit) error
	Update(ctx context.Context, permit *Permit) error
}

// PermitUseCase defines the inbound port for business orchestrations.
type PermitUseCase interface {
	GetPermit(ctx context.Context, id string) (*Permit, error)
	RequestPermit(ctx context.Context, workerID, zoneID string, duration time.Duration) (*Permit, error)
	ApprovePermit(ctx context.Context, id string, supervisorID string) error
}
