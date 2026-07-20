package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	searchDomain "prahari/services/incident/internal/domain/search"
	incidentDomain "prahari/services/incident/internal/domain/incident"
)

// SearchStore implements the search persistence adapter against PostgreSQL.
// It builds dynamic WHERE clauses from search criteria using parameterized queries.
type SearchStore struct {
	db *sql.DB
}

// NewSearchStore constructs a SearchStore.
func NewSearchStore(db *sql.DB) *SearchStore {
	return &SearchStore{db: db}
}

// Search executes a dynamic parameterized query against the incidents table.
func (s *SearchStore) Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*incidentDomain.Incident, int, error) {
	var conditions []string
	var args []interface{}
	argIdx := 1

	conditions = append(conditions, "is_deleted = false")

	if criteria.IncidentNumber != "" {
		conditions = append(conditions, fmt.Sprintf("incident_number = $%d", argIdx))
		args = append(args, criteria.IncidentNumber)
		argIdx++
	}
	if criteria.Category != "" {
		conditions = append(conditions, fmt.Sprintf("category_id = $%d", argIdx))
		args = append(args, criteria.Category)
		argIdx++
	}
	if criteria.Severity != "" {
		conditions = append(conditions, fmt.Sprintf("severity_level = $%d", argIdx))
		args = append(args, criteria.Severity)
		argIdx++
	}
	if criteria.Status != "" {
		conditions = append(conditions, fmt.Sprintf("status_code = $%d", argIdx))
		args = append(args, criteria.Status)
		argIdx++
	}
	if criteria.Type != "" {
		conditions = append(conditions, fmt.Sprintf("type = $%d", argIdx))
		args = append(args, criteria.Type)
		argIdx++
	}
	if criteria.ReporterID != "" {
		conditions = append(conditions, fmt.Sprintf("reporter_id = $%d", argIdx))
		args = append(args, criteria.ReporterID)
		argIdx++
	}
	if criteria.AssigneeID != "" {
		conditions = append(conditions, fmt.Sprintf("assignee_id = $%d", argIdx))
		args = append(args, criteria.AssigneeID)
		argIdx++
	}
	if criteria.Department != "" {
		conditions = append(conditions, fmt.Sprintf("department_id = $%d", argIdx))
		args = append(args, criteria.Department)
		argIdx++
	}
	if criteria.Location != "" {
		conditions = append(conditions, fmt.Sprintf("location_id = $%d", argIdx))
		args = append(args, criteria.Location)
		argIdx++
	}
	if criteria.HasDateRange() {
		conditions = append(conditions, fmt.Sprintf("occurred_at >= $%d", argIdx))
		args = append(args, criteria.DateFrom)
		argIdx++
		conditions = append(conditions, fmt.Sprintf("occurred_at <= $%d", argIdx))
		args = append(args, criteria.DateTo)
		argIdx++
	}
	if criteria.HasFreeText() {
		conditions = append(conditions, fmt.Sprintf("search_vector @@ plainto_tsquery('english', $%d)", argIdx))
		args = append(args, criteria.FreeText)
		argIdx++
	}

	whereClause := strings.Join(conditions, " AND ")

	// Count total matches
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM incidents WHERE %s", whereClause)
	var totalCount int
	if err := s.db.QueryRowContext(ctx, countQuery, args...).Scan(&totalCount); err != nil {
		return nil, 0, fmt.Errorf("postgres: failed to count search results: %w", err)
	}

	// Validate sort column to prevent injection
	allowedSorts := map[string]bool{
		"created_at": true, "occurred_at": true, "severity_level": true,
		"status_code": true, "incident_number": true, "title": true,
	}
	sortBy := "created_at"
	if allowedSorts[criteria.SortBy] {
		sortBy = criteria.SortBy
	}
	sortOrder := "DESC"
	if criteria.SortOrder == "ASC" {
		sortOrder = "ASC"
	}

	// Fetch results with pagination
	dataQuery := fmt.Sprintf(`
		SELECT id, incident_number, title, description, type, category_id,
			severity_level, priority_level, status_code, reporter_id,
			assignee_id, department_id, location_id, location_detail,
			occurred_at, reported_at, resolved_at, closed_at, created_at, updated_at
		FROM incidents WHERE %s
		ORDER BY %s %s
		LIMIT $%d OFFSET $%d`,
		whereClause, sortBy, sortOrder, argIdx, argIdx+1)

	args = append(args, criteria.PageSize, criteria.Offset())

	rows, err := s.db.QueryContext(ctx, dataQuery, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("postgres: failed to execute search query: %w", err)
	}
	defer rows.Close()

	var incidents []*incidentDomain.Incident
	for rows.Next() {
		inc := &incidentDomain.Incident{}
		if err := rows.Scan(
			&inc.ID, &inc.IncidentNumber, &inc.Title, &inc.Description, &inc.Type,
			&inc.CategoryID, &inc.SeverityLevel, &inc.PriorityLevel, &inc.StatusCode,
			&inc.ReporterID, &inc.AssigneeID, &inc.DepartmentID, &inc.LocationID,
			&inc.LocationDetail, &inc.OccurredAt, &inc.ReportedAt, &inc.ResolvedAt,
			&inc.ClosedAt, &inc.CreatedAt, &inc.UpdatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("postgres: failed to scan search result: %w", err)
		}
		incidents = append(incidents, inc)
	}

	return incidents, totalCount, nil
}
