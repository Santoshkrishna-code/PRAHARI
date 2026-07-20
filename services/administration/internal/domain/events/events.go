package events

import "time"

const (
	EventTenantCreated       = "tenant.created"
	EventOrganizationCreated = "organization.created"
	EventPlantCreated        = "plant.created"
	EventDepartmentCreated   = "department.created"
	EventConfigurationUpdated = "configuration.updated"
	EventFeatureflagChanged  = "featureflag.changed"
	EventLicenseUpdated      = "license.updated"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	TenantID  string    `json:"tenant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
