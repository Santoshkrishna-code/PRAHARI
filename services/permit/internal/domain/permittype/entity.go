package permittype

import (
	"encoding/json"
	"errors"
)

// Code defines predefined permit classification tags.
type Code string

const (
	CodeHotWork          Code = "HOT_WORK"
	CodeColdWork          Code = "COLD_WORK"
	CodeConfinedSpace     Code = "CONFINED_SPACE"
	CodeElectrical        Code = "ELECTRICAL"
	CodeMechanical        Code = "MECHANICAL"
	CodeExcavation        Code = "EXCAVATION"
	CodeWorkingAtHeight   Code = "WORKING_AT_HEIGHT"
	CodeRadiation         Code = "RADIATION"
	CodeChemical          Code = "CHEMICAL"
	CodeLineBreaking      Code = "LINE_BREAKING"
	CodeIsolation         Code = "ISOLATION"
	CodeMaintenance       Code = "MAINTENANCE"
)

// PermitType holds template and configuration values for a category of permit.
type PermitType struct {
	ID                  string          `json:"id" db:"id"`
	Code                Code            `json:"code" db:"code"`
	Name                string          `json:"name" db:"name"`
	Description         string          `json:"description" db:"description"`
	DefaultDurationHours int             `json:"default_duration_hours" db:"default_duration_hours"`
	Preconditions       json.RawMessage `json:"preconditions" db:"preconditions"` // JSON array of required flags (e.g. gas_test_required)
	IsActive            bool            `json:"is_active" db:"is_active"`
}

// Validate checks domain invariants for PermitType.
func (pt *PermitType) Validate() error {
	if pt.Name == "" {
		return errors.New("permit type name is required")
	}
	if pt.Code == "" {
		return errors.New("permit type code is required")
	}
	if pt.DefaultDurationHours <= 0 {
		return errors.New("default duration must be greater than zero hours")
	}
	return nil
}
