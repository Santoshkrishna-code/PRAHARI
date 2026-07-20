package export

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"

	assetDomain "prahari/services/asset/internal/domain/asset"
	searchDomain "prahari/services/asset/internal/domain/search"
)

// SearchRepository lists matches.
type SearchRepository interface {
	Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*assetDomain.Asset, int, error)
}

// Service writes raw binary streams.
type Service struct {
	repo SearchRepository
}

// NewService instantiates Export Service.
func NewService(repo SearchRepository) *Service {
	return &Service{repo: repo}
}

// ExportCSV streams a CSV array.
func (s *Service) ExportCSV(ctx context.Context, criteria *searchDomain.Criteria) ([]byte, error) {
	criteria.Normalize()
	criteria.PageSize = 10000

	assets, _, err := s.repo.Search(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve assets for export: %w", err)
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	header := []string{"Asset Number", "Name", "Serial Number", "Lifecycle", "Operational", "Criticality", "Health"}
	_ = writer.Write(header)

	for _, a := range assets {
		row := []string{
			a.AssetNumber,
			a.Name,
			a.SerialNumber,
			a.LifecycleStatus,
			string(a.OperationalStatus),
			string(a.CriticalityCode),
			fmt.Sprintf("%.1f%%", a.HealthScore),
		}
		_ = writer.Write(row)
	}

	writer.Flush()
	return buf.Bytes(), nil
}

// ExportPDF prints layout summary text bytes.
func (s *Service) ExportPDF(ctx context.Context, assetID string) ([]byte, error) {
	criteria := &searchDomain.Criteria{
		AssetNumber: assetID,
		PageSize:    1,
	}
	criteria.Normalize()

	assets, _, err := s.repo.Search(ctx, criteria)
	if err != nil || len(assets) == 0 {
		return nil, fmt.Errorf("failed to retrieve asset: %w", err)
	}

	a := assets[0]

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("ASSET PROFILE SPECIFICATION SHEET\n"))
	buf.WriteString(fmt.Sprintf("=================================\n\n"))
	buf.WriteString(fmt.Sprintf("Asset Number: %s\n", a.AssetNumber))
	buf.WriteString(fmt.Sprintf("Name: %s\n", a.Name))
	buf.WriteString(fmt.Sprintf("Serial Number: %s\n", a.SerialNumber))
	buf.WriteString(fmt.Sprintf("Lifecycle Code: %s\n", a.LifecycleStatus))
	buf.WriteString(fmt.Sprintf("Operational Status: %s\n", a.OperationalStatus))
	buf.WriteString(fmt.Sprintf("Criticality Band: %s\n", a.CriticalityCode))
	buf.WriteString(fmt.Sprintf("Health Index Rating: %.2f%%\n\n", a.HealthScore))
	buf.WriteString(fmt.Sprintf("Description:\n%s\n", a.Description))

	return buf.Bytes(), nil
}
