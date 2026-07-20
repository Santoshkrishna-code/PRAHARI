package assettype

import (
	"errors"
)

// AssetType defines predefined templates of industrial equipment types.
type AssetType struct {
	ID          string `json:"id" db:"id"`
	CategoryID  string `json:"category_id" db:"category_id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	IsActive    bool   `json:"is_active" db:"is_active"`
}

// Validate checks domain invariants.
func (at *AssetType) Validate() error {
	if at.Name == "" {
		return errors.New("asset type name is required")
	}
	if at.CategoryID == "" {
		return errors.New("category ID is required")
	}
	return nil
}
