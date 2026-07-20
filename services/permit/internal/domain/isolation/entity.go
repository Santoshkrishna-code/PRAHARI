package isolation

import (
	"errors"
	"time"
)

// Type defines energy isolation systems.
type Type string

const (
	TypeElectrical Type = "ELECTRICAL"
	TypeMechanical Type = "MECHANICAL"
	TypeProcess    Type = "PROCESS"
	TypePneumatic  Type = "PNEUMATIC"
	TypeHydraulic  Type = "HYDRAULIC"
)

// Status defines Lock-Out/Tag-Out transition stages.
type Status string

const (
	StatusApplied  Status = "APPLIED"
	StatusVerified Status = "VERIFIED"
	StatusRemoved  Status = "REMOVED"
)

// Isolation records a LOTO point application.
type Isolation struct {
	ID                   string     `json:"id" db:"id"`
	PermitID             string     `json:"permit_id" db:"permit_id"`
	IsolationType        Type       `json:"isolation_type" db:"isolation_type"`
	EquipmentID          string     `json:"equipment_id" db:"equipment_id"`
	EquipmentDescription string     `json:"equipment_description" db:"equipment_description"`
	IsolationPoint       string     `json:"isolation_point" db:"isolation_point"`
	LockNumber           string     `json:"lock_number" db:"lock_number"`
	TagNumber            string     `json:"tag_number" db:"tag_number"`
	IsolatedBy           string     `json:"isolated_by" db:"isolated_by"`
	IsolatedAt           time.Time  `json:"isolated_at" db:"isolated_at"`
	VerifiedBy           string     `json:"verified_by,omitempty" db:"verified_by"`
	VerifiedAt           *time.Time `json:"verified_at,omitempty" db:"verified_at"`
	RemovedBy            string     `json:"removed_by,omitempty" db:"removed_by"`
	RemovedAt            *time.Time `json:"removed_at,omitempty" db:"removed_at"`
	Status               Status     `json:"status" db:"status"`
}

// Validate checks domain invariants for Isolation.
func (i *Isolation) Validate() error {
	if i.PermitID == "" {
		return errors.New("permit ID is required for isolation")
	}
	if i.EquipmentID == "" {
		return errors.New("equipment ID is required")
	}
	if i.IsolationPoint == "" {
		return errors.New("isolation point identifier is required")
	}
	if i.LockNumber == "" {
		return errors.New("lock number is required")
	}
	return nil
}

// Verify registers third-party verification details on this LOTO lock.
func (i *Isolation) Verify(verifier string) {
	now := time.Now()
	i.VerifiedBy = verifier
	i.VerifiedAt = &now
	i.Status = StatusVerified
}

// Remove registers release of isolation state.
func (i *Isolation) Remove(remover string) {
	now := time.Now()
	i.RemovedBy = remover
	i.RemovedAt = &now
	i.Status = StatusRemoved
}

// IsVerified checks if verification signatures are complete.
func (i *Isolation) IsVerified() bool {
	return i.Status == StatusVerified
}
