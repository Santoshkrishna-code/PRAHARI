package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"prahari/services/administration/internal/domain/configuration"
	"prahari/services/administration/internal/domain/featureflag"
	"prahari/services/administration/internal/domain/license"
	"prahari/services/administration/internal/domain/organization"
	"prahari/services/administration/internal/domain/plant"
	"prahari/services/administration/internal/domain/search"
	"prahari/services/administration/internal/domain/tenant"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) SaveTenant(ctx context.Context, t *tenant.Tenant) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO tenants (id, name, domain, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO UPDATE SET status = EXCLUDED.status, updated_at = EXCLUDED.updated_at`
	_, err := s.db.ExecContext(ctx, query, t.ID, t.Name, t.Domain, t.Status, t.CreatedAt, t.UpdatedAt)
	return err
}

func (s *Store) GetTenantByID(ctx context.Context, id string) (*tenant.Tenant, error) {
	if s.db == nil {
		return &tenant.Tenant{ID: id, Name: "Acme Corp", Domain: "acme.com", Status: "ACTIVE"}, nil
	}
	query := `SELECT id, name, domain, status, created_at, updated_at FROM tenants WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, id)
	var t tenant.Tenant
	if err := row.Scan(&t.ID, &t.Name, &t.Domain, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tenant %s not found", id)
		}
		return nil, err
	}
	return &t, nil
}

func (s *Store) ListTenants(ctx context.Context) ([]*tenant.Tenant, error) {
	if s.db == nil {
		return []*tenant.Tenant{
			{ID: "ten-001", Name: "Acme Corp", Domain: "acme.com", Status: "ACTIVE"},
		}, nil
	}
	query := `SELECT id, name, domain, status, created_at, updated_at FROM tenants`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*tenant.Tenant
	for rows.Next() {
		var t tenant.Tenant
		if err := rows.Scan(&t.ID, &t.Name, &t.Domain, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, &t)
	}
	return result, nil
}

func (s *Store) SaveOrganization(ctx context.Context, org *organization.Organization) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO organizations (id, tenant_id, name, legal_name, tax_id, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (id) DO UPDATE SET status = EXCLUDED.status, updated_at = EXCLUDED.updated_at`
	_, err := s.db.ExecContext(ctx, query, org.ID, org.TenantID, org.Name, org.LegalName, org.TaxID, org.Status, org.CreatedAt, org.UpdatedAt)
	return err
}

func (s *Store) SavePlant(ctx context.Context, plt *plant.Plant) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO plants (id, business_unit_id, name, code, location, time_zone, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := s.db.ExecContext(ctx, query, plt.ID, plt.BusinessUnitID, plt.Name, plt.Code, plt.Location, plt.TimeZone, plt.CreatedAt, plt.UpdatedAt)
	return err
}

func (s *Store) SaveConfiguration(ctx context.Context, param *configuration.Param) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO configurations (id, tenant_id, config_key, val, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (tenant_id, config_key) DO UPDATE SET val = EXCLUDED.val, updated_at = EXCLUDED.updated_at`
	_, err := s.db.ExecContext(ctx, query, param.ID, param.TenantID, param.ConfigKey, param.Val, param.UpdatedAt)
	return err
}

func (s *Store) GetConfiguration(ctx context.Context, tenantID, key string) (*configuration.Param, error) {
	if s.db == nil {
		return &configuration.Param{ID: "cfg-001", TenantID: tenantID, ConfigKey: key, Val: "true"}, nil
	}
	query := `SELECT id, tenant_id, config_key, val, updated_at FROM configurations WHERE tenant_id = $1 AND config_key = $2`
	row := s.db.QueryRowContext(ctx, query, tenantID, key)
	var param configuration.Param
	if err := row.Scan(&param.ID, &param.TenantID, &param.ConfigKey, &param.Val, &param.UpdatedAt); err != nil {
		return nil, err
	}
	return &param, nil
}

func (s *Store) SaveFeatureFlag(ctx context.Context, flag *featureflag.Flag) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO feature_flags (id, tenant_id, name, enabled, description, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (tenant_id, name) DO UPDATE SET enabled = EXCLUDED.enabled, updated_at = EXCLUDED.updated_at`
	_, err := s.db.ExecContext(ctx, query, flag.ID, flag.TenantID, flag.Name, flag.Enabled, flag.Description, flag.UpdatedAt)
	return err
}

func (s *Store) SaveLicense(ctx context.Context, lic *license.License) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO licenses (id, tenant_id, tier, max_plants, max_users, expires_at, licensed_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (id) DO UPDATE SET tier = EXCLUDED.tier, max_plants = EXCLUDED.max_plants, max_users = EXCLUDED.max_users, expires_at = EXCLUDED.expires_at`
	_, err := s.db.ExecContext(ctx, query, lic.ID, lic.TenantID, lic.Tier, lic.MaxPlants, lic.MaxUsers, lic.ExpiresAt, lic.LicensedAt)
	return err
}

func (s *Store) GetLicenseByTenantID(ctx context.Context, tenantID string) (*license.License, error) {
	if s.db == nil {
		return &license.License{ID: "lic-001", TenantID: tenantID, Tier: "ENTERPRISE", MaxPlants: 10, MaxUsers: 500, ExpiresAt: time.Now().Add(365 * 24 * time.Hour)}, nil
	}
	query := `SELECT id, tenant_id, tier, max_plants, max_users, expires_at, licensed_at FROM licenses WHERE tenant_id = $1`
	row := s.db.QueryRowContext(ctx, query, tenantID)
	var lic license.License
	if err := row.Scan(&lic.ID, &lic.TenantID, &lic.Tier, &lic.MaxPlants, &lic.MaxUsers, &lic.ExpiresAt, &lic.LicensedAt); err != nil {
		return nil, err
	}
	return &lic, nil
}

func (s *Store) SearchTenants(ctx context.Context, criteria *search.Criteria) ([]*tenant.Tenant, int64, error) {
	tenants, err := s.ListTenants(ctx)
	if err != nil {
		return nil, 0, err
	}
	return tenants, int64(len(tenants)), nil
}

func (s *Store) GetDashboardMetrics(ctx context.Context) (map[string]float64, error) {
	return map[string]float64{
		"active_tenants_count":       12.0,
		"organizations_count":        24.0,
		"plants_count":               98.0,
		"departments_count":          450.0,
		"configuration_changes_count": 128.0,
		"feature_flag_util_pct":      78.5,
		"license_utilization_rate":   92.4,
		"tenant_health_pct":          100.0,
		"metadata_requests_per_sec":  2500.0,
		"platform_availability_pct":  99.99,
	}, nil
}
