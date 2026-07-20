package assetcategory

import (
	"errors"
)

// AssetCategory defines general category taxonomies.
type AssetCategory struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	IsActive    bool   `json:"is_active" db:"is_active"`
}

// Validate checks domain invariants.
func (ac *AssetCategory) Validate() error {
	if ac.Name == "" {
		return errors.New("asset category name is required")
	}
	return nil
}
