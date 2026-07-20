package postgres

import (
	"context"
	"database/sql"

	gasDomain "prahari/services/permit/internal/domain/gastest"
)

// GasStore implements atmospheric readings persistent store.
type GasStore struct {
	db *sql.DB
}

// NewGasStore instantiates GasStore.
func NewGasStore(db *sql.DB) *GasStore {
	return &GasStore{db: db}
}

// Create inserts gas test result entries.
func (s *GasStore) Create(ctx context.Context, gt *gasDomain.GasTest) error {
	query := `INSERT INTO permit_gas_tests (id, permit_id, gas_type, reading_value, unit, acceptable_min, acceptable_max, is_passed, tested_by, tested_at, equipment_calibration_date, next_test_due)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err := s.db.ExecContext(ctx, query,
		gt.ID, gt.PermitID, gt.GasType, gt.ReadingValue, gt.Unit, gt.AcceptableMin, gt.AcceptableMax, gt.IsPassed, gt.TestedBy, gt.TestedAt, gt.EquipmentCalibrationDate, gt.NextTestDue,
	)
	return err
}

// FindByPermitID returns values recorded on a permit area.
func (s *GasStore) FindByPermitID(ctx context.Context, permitID string) ([]*gasDomain.GasTest, error) {
	query := `SELECT id, permit_id, gas_type, reading_value, unit, acceptable_min, acceptable_max, is_passed, tested_by, tested_at, equipment_calibration_date, next_test_due
		FROM permit_gas_tests WHERE permit_id = $1`
	rows, err := s.db.QueryContext(ctx, query, permitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tests []*gasDomain.GasTest
	for rows.Next() {
		gt := &gasDomain.GasTest{}
		err = rows.Scan(&gt.ID, &gt.PermitID, &gt.GasType, &gt.ReadingValue, &gt.Unit, &gt.AcceptableMin, &gt.AcceptableMax, &gt.IsPassed, &gt.TestedBy, &gt.TestedAt, &gt.EquipmentCalibrationDate, &gt.NextTestDue)
		if err != nil {
			return nil, err
		}
		tests = append(tests, gt)
	}
	return tests, nil
}
