package identifier

import (
	"errors"
)

// Type defines tags scanning categories.
type Type string

const (
	TypeQRCode  Type = "QR_CODE"
	TypeBarcode Type = "BARCODE"
	TypeRFID    Type = "RFID"
	TypeNFC     Type = "NFC"
)

// AssetIdentifier tracks scanning codes mapped to assets.
type AssetIdentifier struct {
	ID      string `json:"id" db:"id"`
	AssetID string `json:"asset_id" db:"asset_id"`
	Type    Type   `json:"type" db:"type"`
	Value   string `json:"value" db:"value"`
}

// Validate checks domain invariants.
func (ai *AssetIdentifier) Validate() error {
	if ai.AssetID == "" {
		return errors.New("asset ID is required for identifier")
	}
	if ai.Value == "" {
		return errors.New("identifier tag value is required")
	}
	return nil
}
