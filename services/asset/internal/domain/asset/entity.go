package asset

import (
	"errors"
	"time"
)

// CriticalityLevel defines priority severity.
type CriticalityLevel string

const (
	CritLow      CriticalityLevel = "LOW"
	CritMedium   CriticalityLevel = "MEDIUM"
	CritHigh     CriticalityLevel = "HIGH"
	CritCritical CriticalityLevel = "CRITICAL"
)

// OperationalStatus defines real-time status.
type OperationalStatus string

const (
	OpRunning   OperationalStatus = "RUNNING"
	OpStopped   OperationalStatus = "STOPPED"
	OpIdle      OperationalStatus = "IDLE"
	OpFaulted   OperationalStatus = "FAULTED"
	OpEmergency OperationalStatus = "EMERGENCY_STOP"
	OpUnavail   OperationalStatus = "UNAVAILABLE"
)

// Asset represents the core aggregate root of the Asset Management domain.
type Asset struct {
	ID                 string            `json:"id" db:"id"`
	AssetNumber        string            `json:"asset_number" db:"asset_number"`
	Name               string            `json:"name" db:"name"`
	Description        string            `json:"description" db:"description"`
	SerialNumber       string            `json:"serial_number" db:"serial_number"`
	LifecycleStatus    string            `json:"lifecycle_status" db:"lifecycle_status"`
	OperationalStatus  OperationalStatus `json:"operational_status" db:"operational_status"`
	CriticalityCode    CriticalityLevel  `json:"criticality_code" db:"criticality_code"`
	DepartmentID       string            `json:"department_id" db:"department_id"`
	LocationID         string            `json:"location_id" db:"location_id"`
	CategoryID         string            `json:"category_id" db:"category_id"`
	TypeID             string            `json:"type_id" db:"type_id"`
	ManufacturerID     string            `json:"manufacturer_id" db:"manufacturer_id"`
	ModelNumber        string            `json:"model_number" db:"model_number"`
	PurchaseDate       time.Time         `json:"purchase_date" db:"purchase_date"`
	InstallationDate   *time.Time        `json:"installation_date,omitempty" db:"installation_date"`
	LastMaintenanceDate *time.Time       `json:"last_maintenance_date,omitempty" db:"last_maintenance_date"`
	HealthScore        float64           `json:"health_score" db:"health_score"`               // 0% - 100%
	ConditionScore     float64           `json:"condition_score" db:"condition_score"`         // physical condition
	RemainingUsefulLife float64          `json:"remaining_useful_life" db:"remaining_useful_life"` // in hours
	FailureProbability float64           `json:"failure_probability" db:"failure_probability"` // 0.0 - 1.0
	CreatedAt          time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time         `json:"updated_at" db:"updated_at"`
	IsDeleted          bool              `json:"is_deleted" db:"is_deleted"`
}

// Validate checks domain invariants for Asset.
func (a *Asset) Validate() error {
	if a.Name == "" {
		return errors.New("asset name is required")
	}
	if len(a.Name) > 200 {
		return errors.New("asset name must not exceed 200 characters")
	}
	if a.AssetNumber == "" {
		return errors.New("asset number is required")
	}
	if a.DepartmentID == "" {
		return errors.New("department ID is required")
	}
	if a.LocationID == "" {
		return errors.New("location ID is required")
	}
	return nil
}
