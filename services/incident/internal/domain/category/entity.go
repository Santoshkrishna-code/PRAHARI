package category

import (
	"errors"
)

// Category represents an incident classification type within a hierarchical taxonomy.
// Categories can be nested via ParentID to support multi-level classification
// (e.g., "Equipment" → "Electrical" → "Arc Flash").
type Category struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	ParentID    string `json:"parent_id,omitempty" db:"parent_id"`
	SortOrder   int    `json:"sort_order" db:"sort_order"`
	IsActive    bool   `json:"is_active" db:"is_active"`
}

// DefaultCategories defines the standard safety incident classification taxonomy.
var DefaultCategories = []Category{
	{Name: "Fire", Description: "Fire-related incidents including explosions and thermal events"},
	{Name: "Chemical Spill", Description: "Hazardous chemical releases and contamination events"},
	{Name: "Electrical", Description: "Electrical hazards including arc flash and electrocution"},
	{Name: "Fall", Description: "Falls from height or same-level slips, trips, and falls"},
	{Name: "Equipment Failure", Description: "Mechanical or equipment malfunction events"},
	{Name: "Environmental", Description: "Environmental contamination and ecological impact events"},
	{Name: "Structural", Description: "Structural integrity failures including collapses"},
	{Name: "Transportation", Description: "Vehicle and transportation-related incidents"},
	{Name: "Biological", Description: "Biological hazard exposures and infectious events"},
	{Name: "Ergonomic", Description: "Repetitive strain and ergonomic-related injuries"},
	{Name: "Other", Description: "Incidents not covered by standard categories"},
}

// Validate enforces domain invariants on the category aggregate.
func (c *Category) Validate() error {
	if c.Name == "" {
		return errors.New("category name is required")
	}
	if len(c.Name) > 200 {
		return errors.New("category name must not exceed 200 characters")
	}
	return nil
}

// IsRoot returns true if this category has no parent (is a top-level classification).
func (c *Category) IsRoot() bool {
	return c.ParentID == ""
}
