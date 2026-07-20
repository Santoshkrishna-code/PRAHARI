package assethierarchy

import (
	"errors"
)

// RelationType defines structure links categories.
type RelationType string

const (
	TypeComponentOf   RelationType = "COMPONENT_OF"
	TypeSubAssemblyOf RelationType = "SUB_ASSEMBLY_OF"
)

// AssetHierarchy models parent-child assets tree structures.
type AssetHierarchy struct {
	ID               string       `json:"id" db:"id"`
	ParentAssetID    string       `json:"parent_asset_id" db:"parent_asset_id"`
	ChildAssetID     string       `json:"child_asset_id" db:"child_asset_id"`
	RelationshipType RelationType `json:"relationship_type" db:"relationship_type"`
}

// Validate checks domain invariants.
func (h *AssetHierarchy) Validate() error {
	if h.ParentAssetID == "" {
		return errors.New("parent asset ID is required")
	}
	if h.ChildAssetID == "" {
		return errors.New("child asset ID is required")
	}
	return nil
}
