package visit

import "time"

// Visit represents a scheduled or active visit instance.
type Visit struct {
	ID          string    `json:"id"`
	VisitorID   string    `json:"visitor_id"`
	HostID      string    `json:"host_id"`
	PlantID     string    `json:"plant_id"`
	Purpose     string    `json:"purpose"`
	ScheduledIn time.Time `json:"scheduled_in"`
	ScheduledOut time.Time `json:"scheduled_out"`
	Status      string    `json:"status"` // Scheduled, Host Approval, Security Verification, Gate Pass Issued, Checked In, On Site, Checked Out, Closed, Rejected, Cancelled, Blacklisted
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
