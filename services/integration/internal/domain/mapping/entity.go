package mapping

import "time"

// FieldMap represents a data mapping schema rule to translate external keys to internal EHS keys.
type FieldMap struct {
	ID          string    `json:"id"`
	ConnectorID string    `json:"connector_id"`
	ExternalKey string    `json:"external_key"`
	InternalKey string    `json:"internal_key"`
	DataType    string    `json:"data_type"` // STRING, INT, FLOAT, BOOLEAN
	UpdatedAt   time.Time `json:"updated_at"`
}
