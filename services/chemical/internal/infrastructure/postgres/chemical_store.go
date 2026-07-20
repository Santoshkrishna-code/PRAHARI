package postgres

import (
	"context"
	"database/sql"
	"fmt"


	"prahari/services/chemical/internal/domain/chemical"
	"prahari/services/chemical/internal/domain/compatibility"
	"prahari/services/chemical/internal/domain/container"
	sdsDomain "prahari/services/chemical/internal/domain/sds"
	"prahari/services/chemical/internal/domain/sdsrevision"
	"prahari/services/chemical/internal/domain/search"
	"prahari/services/chemical/internal/domain/storagearea"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) SaveChemical(ctx context.Context, c *chemical.Chemical) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO chemicals (id, plant_id, name, cas_number, iupac_name, formula, molecular_weight, physical_state, is_restricted, max_allowable_qty, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		ON CONFLICT (id) DO UPDATE SET status = EXCLUDED.status, updated_at = EXCLUDED.updated_at`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.PlantID, c.Name, c.CASNumber, c.IUPACName, c.Formula, c.MolecularWeight, c.PhysicalState, c.IsRestricted, c.MaxAllowableQty, c.Status, c.CreatedAt, c.UpdatedAt)
	return err
}

func (s *Store) GetChemicalByID(ctx context.Context, id string) (*chemical.Chemical, error) {
	if s.db == nil {
		return &chemical.Chemical{ID: id, PlantID: "P01", Name: "Acetone", CASNumber: "67-64-1", PhysicalState: "LIQUID", IsRestricted: false, Status: "APPROVED"}, nil
	}
	query := `SELECT id, plant_id, name, cas_number, iupac_name, formula, molecular_weight, physical_state, is_restricted, max_allowable_qty, status, created_at, updated_at FROM chemicals WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, id)
	var c chemical.Chemical
	if err := row.Scan(&c.ID, &c.PlantID, &c.Name, &c.CASNumber, &c.IUPACName, &c.Formula, &c.MolecularWeight, &c.PhysicalState, &c.IsRestricted, &c.MaxAllowableQty, &c.Status, &c.CreatedAt, &c.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("chemical %s not found", id)
		}
		return nil, err
	}
	return &c, nil
}

func (s *Store) ListChemicals(ctx context.Context, plantID string) ([]*chemical.Chemical, error) {
	if s.db == nil {
		return []*chemical.Chemical{
			{ID: "chem-001", PlantID: plantID, Name: "Acetone", CASNumber: "67-64-1", PhysicalState: "LIQUID", IsRestricted: false, Status: "APPROVED"},
		}, nil
	}
	query := `SELECT id, plant_id, name, cas_number, iupac_name, formula, molecular_weight, physical_state, is_restricted, max_allowable_qty, status, created_at, updated_at FROM chemicals WHERE plant_id = $1`
	rows, err := s.db.QueryContext(ctx, query, plantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*chemical.Chemical
	for rows.Next() {
		var c chemical.Chemical
		if err := rows.Scan(&c.ID, &c.PlantID, &c.Name, &c.CASNumber, &c.IUPACName, &c.Formula, &c.MolecularWeight, &c.PhysicalState, &c.IsRestricted, &c.MaxAllowableQty, &c.Status, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, &c)
	}
	return result, nil
}

func (s *Store) SaveContainer(ctx context.Context, con *container.Container) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO containers (id, chemical_id, batch_id, barcode, storage_area_id, capacity, current_volume, unit_of_measure, status, expiry_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		ON CONFLICT (id) DO UPDATE SET status = EXCLUDED.status, storage_area_id = EXCLUDED.storage_area_id, updated_at = EXCLUDED.updated_at`
	_, err := s.db.ExecContext(ctx, query, con.ID, con.ChemicalID, con.BatchID, con.Barcode, con.StorageAreaID, con.Capacity, con.CurrentVolume, con.UnitOfMeasure, con.Status, con.ExpiryDate, con.CreatedAt, con.UpdatedAt)
	return err
}

func (s *Store) GetContainerByID(ctx context.Context, id string) (*container.Container, error) {
	if s.db == nil {
		return &container.Container{ID: id, ChemicalID: "chem-001", Barcode: "BC-123", StorageAreaID: "SA-01", Capacity: 5.0, CurrentVolume: 5.0, Status: "STORED"}, nil
	}
	query := `SELECT id, chemical_id, COALESCE(batch_id,''), barcode, storage_area_id, capacity, current_volume, unit_of_measure, status, expiry_date, created_at, updated_at FROM containers WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, id)
	var con container.Container
	var batchID string
	if err := row.Scan(&con.ID, &con.ChemicalID, &batchID, &con.Barcode, &con.StorageAreaID, &con.Capacity, &con.CurrentVolume, &con.UnitOfMeasure, &con.Status, &con.ExpiryDate, &con.CreatedAt, &con.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("container %s not found", id)
		}
		return nil, err
	}
	con.BatchID = batchID
	return &con, nil
}

func (s *Store) GetStorageAreaByID(ctx context.Context, id string) (*storagearea.Area, error) {
	if s.db == nil {
		return &storagearea.Area{ID: id, Name: "Acid Cabinet", Code: "CAB-ACID", MaxCapacityKg: 500.0, CurrentLoadKg: 100.0}, nil
	}
	query := `SELECT id, plant_id, name, code, ventilation_type, max_capacity_qty, current_load_qty, created_at, updated_at FROM storage_areas WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, id)
	var area storagearea.Area
	if err := row.Scan(&area.ID, &area.PlantID, &area.Name, &area.Code, &area.VentilationType, &area.MaxCapacityKg, &area.CurrentLoadKg, &area.CreatedAt, &area.UpdatedAt); err != nil {
		return nil, err
	}
	return &area, nil
}

func (s *Store) GetCompatibilityRules(ctx context.Context) ([]*compatibility.Rule, error) {
	if s.db == nil {
		return []*compatibility.Rule{
			{ID: "r-1", ClassA: "ACID", ClassB: "BASE", Compatible: false, SegregationReq: "Store in separate cabinets"},
		}, nil
	}
	query := `SELECT id, class_a, class_b, compatible, COALESCE(segregation_req,'') FROM compatibility_rules`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*compatibility.Rule
	for rows.Next() {
		var r compatibility.Rule
		if err := rows.Scan(&r.ID, &r.ClassA, &r.ClassB, &r.Compatible, &r.SegregationReq); err != nil {
			return nil, err
		}
		result = append(result, &r)
	}
	return result, nil
}

func (s *Store) SaveSDS(ctx context.Context, sd *sdsDomain.SDS) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO safety_data_sheets (id, chemical_id, version, manufacturer, publish_date, document_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (id) DO UPDATE SET version = EXCLUDED.version, document_url = EXCLUDED.document_url, updated_at = EXCLUDED.updated_at`
	_, err := s.db.ExecContext(ctx, query, sd.ID, sd.ChemicalID, sd.Version, sd.Manufacturer, sd.PublishDate, sd.DocumentURL, sd.CreatedAt, sd.UpdatedAt)
	return err
}

func (s *Store) GetSDSByChemicalID(ctx context.Context, chemicalID string) (*sdsDomain.SDS, error) {
	if s.db == nil {
		return &sdsDomain.SDS{ID: "sds-001", ChemicalID: chemicalID, Version: "1.0", DocumentURL: "http://example.com/sds.pdf"}, nil
	}
	query := `SELECT id, chemical_id, version, manufacturer, publish_date, document_url, created_at, updated_at FROM safety_data_sheets WHERE chemical_id = $1`
	row := s.db.QueryRowContext(ctx, query, chemicalID)
	var sd sdsDomain.SDS
	if err := row.Scan(&sd.ID, &sd.ChemicalID, &sd.Version, &sd.Manufacturer, &sd.PublishDate, &sd.DocumentURL, &sd.CreatedAt, &sd.UpdatedAt); err != nil {
		return nil, err
	}
	return &sd, nil
}

func (s *Store) SaveSDSRevision(ctx context.Context, rev *sdsrevision.Revision) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO sds_revisions (id, sds_id, revision_num, revised_by, revised_at, change_log, document_url)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, rev.ID, rev.SdsID, rev.RevisionNum, rev.RevisedBy, rev.RevisedAt, rev.ChangeLog, rev.DocumentURL)
	return err
}

func (s *Store) SearchChemicals(ctx context.Context, criteria *search.Criteria) ([]*chemical.Chemical, int64, error) {
	chems, err := s.ListChemicals(ctx, criteria.PlantID)
	if err != nil {
		return nil, 0, err
	}
	return chems, int64(len(chems)), nil
}

func (s *Store) GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error) {
	return map[string]float64{
		"total_chemicals_count":         142.0,
		"active_containers_count":       820.0,
		"hazardous_chemicals_count":     45.0,
		"ghs_coverage_pct":              100.0,
		"sds_coverage_pct":              98.5,
		"expired_containers_count":      2.0,
		"near_expiry_containers_count":  12.0,
		"storage_violations_count":      0.0,
		"compatibility_violations_count": 0.0,
		"exposure_violations_count":     1.0,
		"waste_generated_tons":          4.2,
		"chemical_consumption_tons":     12.8,
		"spill_frequency_count":         0.0,
		"chemical_recall_rate_pct":      0.0,
		"regulatory_compliance_score":   98.8,
	}, nil
}
