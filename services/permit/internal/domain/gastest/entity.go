package gastest

import (
	"errors"
	"time"
)

// GasType defines target toxic/explosive elements.
type GasType string

const (
	GasTypeO2      GasType = "O2"
	GasTypeLEL     GasType = "LEL"
	GasTypeH2S     GasType = "H2S"
	GasTypeCO      GasType = "CO"
	GasTypeSO2     GasType = "SO2"
	GasTypeBenzene GasType = "BENZENE"
)

// GasTest records atmospheric safety samples inside confined areas.
type GasTest struct {
	ID                       string    `json:"id" db:"id"`
	PermitID                 string    `json:"permit_id" db:"permit_id"`
	GasType                  GasType   `json:"gas_type" db:"gas_type"`
	ReadingValue             float64   `json:"reading_value" db:"reading_value"`
	Unit                     string    `json:"unit" db:"unit"`
	AcceptableMin            float64   `json:"acceptable_min" db:"acceptable_min"`
	AcceptableMax            float64   `json:"acceptable_max" db:"acceptable_max"`
	IsPassed                 bool      `json:"is_passed" db:"is_passed"`
	TestedBy                 string    `json:"tested_by" db:"tested_by"`
	TestedAt                 time.Time `json:"tested_at" db:"tested_at"`
	EquipmentCalibrationDate time.Time `json:"equipment_calibration_date" db:"equipment_calibration_date"`
	NextTestDue              time.Time `json:"next_test_due" db:"next_test_due"`
}

// Validate checks domain invariants for GasTest.
func (gt *GasTest) Validate() error {
	if gt.PermitID == "" {
		return errors.New("permit ID is required for gas test")
	}
	if gt.TestedBy == "" {
		return errors.New("testing officer ID is required")
	}
	if gt.TestedAt.IsZero() {
		return errors.New("test timestamp is required")
	}
	return nil
}

// EvaluateResult evaluates if the current reading complies with safe thresholds.
func (gt *GasTest) EvaluateResult() {
	gt.IsPassed = gt.ReadingValue >= gt.AcceptableMin && gt.ReadingValue <= gt.AcceptableMax
}
