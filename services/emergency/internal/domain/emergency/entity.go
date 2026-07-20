package emergency

import "time"

// Category specifies the physical/operational nature of the emergency event.
type Category string

const (
	CategoryFire           Category = "FIRE"
	CategoryExplosion      Category = "EXPLOSION"
	CategoryToxicRelease   Category = "TOXIC_RELEASE"
	CategoryChemicalSpill  Category = "CHEMICAL_SPILL"
	CategoryGasLeak        Category = "GAS_LEAK"
	CategoryMedical        Category = "MEDICAL_EMERGENCY"
	CategoryNaturalDisaster Category = "NATURAL_DISASTER"
	CategoryFlood          Category = "FLOOD"
	CategoryEarthquake     Category = "EARTHQUAKE"
	CategoryCyclone        Category = "CYCLONE"
	CategoryPowerFailure   Category = "POWER_FAILURE"
	CategoryCyber          Category = "CYBER_INCIDENT"
	CategorySecurity       Category = "SECURITY_THREAT"
)

// Emergency represents an active or historical industrial emergency event.
type Emergency struct {
	ID                 string     `json:"id"`
	EmergencyNumber    string     `json:"emergency_number"`
	PlantID            string     `json:"plant_id"`
	UnitID             string     `json:"unit_id"`
	Title              string     `json:"title"`
	Description        string     `json:"description"`
	Category           Category   `json:"category"`
	Severity           string     `json:"severity"`             // TIER_1, TIER_2, TIER_3 (Major Industrial Emergency)
	IncidentID         string     `json:"incident_id,omitempty"` // Triggering incident ID if applicable
	Status             string     `json:"status"`               // Prepared, Detected, Declared, Response Activated, Command Established, Resource Deployment, Evacuation, Stabilized, Recovery, After Action Review, Closed, Cancelled
	CommanderID        string     `json:"commander_id"`
	DeclaredAt         time.Time  `json:"declared_at"`
	CommandEstablishedAt *time.Time `json:"command_established_at,omitempty"`
	StabilizedAt       *time.Time `json:"stabilized_at,omitempty"`
	ClosedAt           *time.Time `json:"closed_at,omitempty"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}
