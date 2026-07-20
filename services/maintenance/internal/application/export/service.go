package export

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"

	maintenanceDomain "prahari/services/maintenance/internal/domain/maintenance"
	searchDomain "prahari/services/maintenance/internal/domain/search"
)

// SearchRepository lists matches.
type SearchRepository interface {
	Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*maintenanceDomain.Maintenance, int, error)
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

	maintenanceList, _, err := s.repo.Search(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve maintenance records for export: %w", err)
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	header := []string{"Maintenance Number", "Asset ID", "Type", "Priority", "Status", "Estimated Cost", "Actual Cost"}
	_ = writer.Write(header)

	for _, m := range maintenanceList {
		row := []string{
			m.MaintenanceNumber,
			m.AssetID,
			m.MaintenanceType,
			string(m.Priority),
			m.StatusCode,
			fmt.Sprintf("$%.2f", m.TotalEstimatedCost),
			fmt.Sprintf("$%.2f", m.TotalActualCost),
		}
		_ = writer.Write(row)
	}

	writer.Flush()
	return buf.Bytes(), nil
}

// ExportPDF prints layout summary text bytes.
func (s *Service) ExportPDF(ctx context.Context, maintenanceID string) ([]byte, error) {
	criteria := &searchDomain.Criteria{
		MaintenanceNumber: maintenanceID,
		PageSize:          1,
	}
	criteria.Normalize()

	maintenanceList, _, err := s.repo.Search(ctx, criteria)
	if err != nil || len(maintenanceList) == 0 {
		return nil, fmt.Errorf("failed to retrieve maintenance record: %w", err)
	}

	m := maintenanceList[0]

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("MAINTENANCE WORK ORDER EXECUTION PROFILE\n"))
	buf.WriteString(fmt.Sprintf("========================================\n\n"))
	buf.WriteString(fmt.Sprintf("Maintenance Number: %s\n", m.MaintenanceNumber))
	buf.WriteString(fmt.Sprintf("Asset ID: %s\n", m.AssetID))
	buf.WriteString(fmt.Sprintf("Type: %s\n", m.MaintenanceType))
	buf.WriteString(fmt.Sprintf("Priority Level: %s\n", m.Priority))
	buf.WriteString(fmt.Sprintf("Status: %s\n", m.StatusCode))
	buf.WriteString(fmt.Sprintf("Estimated Costs: $%.2f\n", m.TotalEstimatedCost))
	buf.WriteString(fmt.Sprintf("Actual Costs: $%.2f\n\n", m.TotalActualCost))
	buf.WriteString(fmt.Sprintf("Description:\n%s\n", m.Description))

	return buf.Bytes(), nil
}
