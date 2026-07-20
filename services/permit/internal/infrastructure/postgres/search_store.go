package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	searchDomain "prahari/services/permit/internal/domain/search"
	permitDomain "prahari/services/permit/internal/domain/permit"
)

// SearchStore implements dyn query builder against permit tables.
type SearchStore struct {
	db *sql.DB
}

// NewSearchStore instantiates SearchStore.
func NewSearchStore(db *sql.DB) *SearchStore {
	return &SearchStore{db: db}
}

// Search queries using parameter arguments dynamically.
func (s *SearchStore) Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*permitDomain.Permit, int, error) {
	var conditions []string
	var args []interface{}
	argIdx := 1

	conditions = append(conditions, "is_deleted = false")

	if criteria.PermitNumber != "" {
		conditions = append(conditions, fmt.Sprintf("permit_number = $%d", argIdx))
		args = append(args, criteria.PermitNumber)
		argIdx++
	}
	if criteria.PermitType != "" {
		conditions = append(conditions, fmt.Sprintf("permit_type_id = $%d", argIdx))
		args = append(args, criteria.PermitType)
		argIdx++
	}
	if criteria.Status != "" {
		conditions = append(conditions, fmt.Sprintf("status_code = $%d", argIdx))
		args = append(args, criteria.Status)
		argIdx++
	}
	if criteria.ApplicantID != "" {
		conditions = append(conditions, fmt.Sprintf("applicant_id = $%d", argIdx))
		args = append(args, criteria.ApplicantID)
		argIdx++
	}
	if criteria.ApproverID != "" {
		conditions = append(conditions, fmt.Sprintf("supervisor_id = $%d", argIdx))
		args = append(args, criteria.ApproverID)
		argIdx++
	}
	if criteria.Department != "" {
		conditions = append(conditions, fmt.Sprintf("department_id = $%d", argIdx))
		args = append(args, criteria.Department)
		argIdx++
	}
	if criteria.Contractor != "" {
		conditions = append(conditions, fmt.Sprintf("contractor_id = $%d", argIdx))
		args = append(args, criteria.Contractor)
		argIdx++
	}
	if criteria.WorkArea != "" {
		conditions = append(conditions, fmt.Sprintf("work_area_id = $%d", argIdx))
		args = append(args, criteria.WorkArea)
		argIdx++
	}
	if criteria.RiskLevel != "" {
		conditions = append(conditions, fmt.Sprintf("risk_level = $%d", argIdx))
		args = append(args, criteria.RiskLevel)
		argIdx++
	}
	if criteria.HasDateRange() {
		conditions = append(conditions, fmt.Sprintf("planned_start_at >= $%d", argIdx))
		args = append(args, criteria.DateFrom)
		argIdx++
		conditions = append(conditions, fmt.Sprintf("planned_start_at <= $%d", argIdx))
		args = append(args, criteria.DateTo)
		argIdx++
	}
	if criteria.FreeText != "" {
		conditions = append(conditions, fmt.Sprintf("search_vector @@ plainto_tsquery('english', $%d)", argIdx))
		args = append(args, criteria.FreeText)
		argIdx++
	}

	whereClause := strings.Join(conditions, " AND ")

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM permits WHERE %s", whereClause)
	var total int
	if err := s.db.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	allowedSorts := map[string]bool{"created_at": true, "planned_start_at": true, "permit_number": true}
	sortBy := "created_at"
	if allowedSorts[criteria.SortBy] {
		sortBy = criteria.SortBy
	}
	sortOrder := "DESC"
	if criteria.SortOrder == "ASC" {
		sortOrder = "ASC"
	}

	dataQuery := fmt.Sprintf(`SELECT id, permit_number, title, description, permit_type_id, status_code, risk_level,
		applicant_id, supervisor_id, issuer_id, receiver_id, department_id, contractor_id,
		work_area_id, work_description, planned_start_at, planned_end_at, actual_start_at,
		actual_end_at, valid_until, extension_count, linked_incident_id, created_at, updated_at
		FROM permits WHERE %s ORDER BY %s %s LIMIT $%d OFFSET $%d`,
		whereClause, sortBy, sortOrder, argIdx, argIdx+1)

	args = append(args, criteria.PageSize, criteria.Offset())

	rows, err := s.db.QueryContext(ctx, dataQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var permits []*permitDomain.Permit
	for rows.Next() {
		p := &permitDomain.Permit{}
		err = rows.Scan(
			&p.ID, &p.PermitNumber, &p.Title, &p.Description, &p.PermitTypeID, &p.StatusCode, &p.RiskLevel,
			&p.ApplicantID, &p.SupervisorID, &p.IssuerID, &p.ReceiverID, &p.DepartmentID, &p.ContractorID,
			&p.WorkAreaID, &p.WorkDescription, &p.PlannedStartAt, &p.PlannedEndAt, &p.ActualStartAt,
			&p.ActualEndAt, &p.ValidUntil, &p.ExtensionCount, &p.LinkedIncidentID, &p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		permits = append(permits, p)
	}

	return permits, total, nil
}
