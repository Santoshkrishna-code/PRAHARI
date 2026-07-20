package costcenter

import "time"

// CostCenter represents a financial mapping division code.
type CostCenter struct {
	ID             string    `json:"id"`
	OrganizationID string    `json:"organization_id"`
	Code           string    `json:"code"`
	Name           string    `json:"name"`
	ManagerID      string    `json:"manager_id"`
	CreatedAt      time.Time `json:"created_at"`
}
