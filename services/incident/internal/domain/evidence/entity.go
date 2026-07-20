package evidence

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// Type classifies the nature of the evidence collected.
type Type string

const (
	TypePhoto            Type = "PHOTO"
	TypeVideo            Type = "VIDEO"
	TypeDocument         Type = "DOCUMENT"
	TypeWitnessStatement Type = "WITNESS_STATEMENT"
)

// ValidTypes enumerates all accepted evidence type classifications.
var ValidTypes = []Type{
	TypePhoto,
	TypeVideo,
	TypeDocument,
	TypeWitnessStatement,
}

// CustodyEntry records a single transfer in the chain of custody.
type CustodyEntry struct {
	HandledBy   string    `json:"handled_by"`
	Action      string    `json:"action"`
	Timestamp   time.Time `json:"timestamp"`
	Description string    `json:"description"`
}

// Evidence represents a piece of evidence collected during an incident investigation.
// It maintains a chain of custody for legal and compliance traceability.
type Evidence struct {
	ID              string          `json:"id" db:"id"`
	IncidentID      string          `json:"incident_id" db:"incident_id"`
	Type            Type            `json:"type" db:"type"`
	Description     string          `json:"description" db:"description"`
	StoragePath     string          `json:"storage_path" db:"storage_path"`
	CollectedBy     string          `json:"collected_by" db:"collected_by"`
	CollectedAt     time.Time       `json:"collected_at" db:"collected_at"`
	ChainOfCustody  json.RawMessage `json:"chain_of_custody" db:"chain_of_custody"`
}

// Validate enforces domain invariants on the evidence aggregate.
func (e *Evidence) Validate() error {
	if e.IncidentID == "" {
		return errors.New("incident ID is required for evidence")
	}
	if !e.isValidType() {
		return fmt.Errorf("invalid evidence type: %s", e.Type)
	}
	if e.Description == "" {
		return errors.New("evidence description is required")
	}
	if e.CollectedBy == "" {
		return errors.New("collector identity is required")
	}
	return nil
}

// AddCustodyEntry appends a new transfer record to the chain of custody.
func (e *Evidence) AddCustodyEntry(entry CustodyEntry) error {
	var chain []CustodyEntry
	if e.ChainOfCustody != nil {
		if err := json.Unmarshal(e.ChainOfCustody, &chain); err != nil {
			return fmt.Errorf("failed to parse chain of custody: %w", err)
		}
	}
	chain = append(chain, entry)
	data, err := json.Marshal(chain)
	if err != nil {
		return fmt.Errorf("failed to serialize chain of custody: %w", err)
	}
	e.ChainOfCustody = data
	return nil
}

// isValidType checks whether the evidence type is among accepted classifications.
func (e *Evidence) isValidType() bool {
	for _, t := range ValidTypes {
		if e.Type == t {
			return true
		}
	}
	return false
}
