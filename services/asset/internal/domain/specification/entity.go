package specification

import (
	"encoding/json"
	"errors"
)

// Specification maps technical details specifications.
type Specification struct {
	ID         string          `json:"id" db:"id"`
	AssetID    string          `json:"asset_id" db:"asset_id"`
	Attributes json.RawMessage `json:"attributes" db:"attributes"` // JSON key-value map of specs
}

// Validate checks domain invariants.
func (s *Specification) Validate() error {
	if s.AssetID == "" {
		return errors.New("asset ID is required for specifications")
	}
	return nil
}
