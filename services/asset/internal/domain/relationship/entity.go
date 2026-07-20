package relationship

import (
	"errors"
)

// DependencyType defines operational dependencies.
type DependencyType string

const (
	TypeDependsOn DependencyType = "DEPENDS_ON"
	TypePoweredBy DependencyType = "POWERED_BY"
)

// Relationship models operational dependencies between separate assets.
type Relationship struct {
	ID             string         `json:"id" db:"id"`
	SourceAssetID  string         `json:"source_asset_id" db:"source_asset_id"`
	TargetAssetID  string         `json:"target_asset_id" db:"target_asset_id"`
	DependencyType DependencyType `json:"dependency_type" db:"dependency_type"`
}

// Validate checks domain invariants.
func (r *Relationship) Validate() error {
	if r.SourceAssetID == "" {
		return errors.New("source asset ID is required")
	}
	if r.TargetAssetID == "" {
		return errors.New("target asset ID is required")
	}
	return nil
}
