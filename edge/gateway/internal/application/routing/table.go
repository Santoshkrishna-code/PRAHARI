package routing

import (
	"errors"
	"strings"

	"prahari/edge/gateway/internal/domain/route"
)

// Table acts as the in-memory route rules database.
type Table struct {
	routes []route.Route
}

// NewTable constructs a Table.
func NewTable(routes []route.Route) *Table {
	return &Table{routes: routes}
}

// Match resolves the matching route rule.
// E.g. matches prefixes path: "/api/v1/incidents/create" -> "/api/v1/incidents/*"
func (t *Table) Match(path string) (*route.Route, error) {
	for _, r := range t.routes {
		cleanPattern := strings.TrimSuffix(r.Path, "*")
		if strings.HasPrefix(path, cleanPattern) {
			return &r, nil
		}
	}
	return nil, errors.New("routing table check: no matching route found")
}
