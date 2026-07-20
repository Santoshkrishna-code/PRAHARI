package independentprotectionlayer

import "time"

// IPL represents an Independent Protection Layer compliant with CCPS LOPA criteria.
type IPL struct {
	ID            string    `json:"id"`
	BarrierID     string    `json:"barrier_id"`
	IPLName       string    `json:"ipl_name"`
	PFDClaimed    float64   `json:"pfd_claimed"` // E.g., 0.01 for 100x risk reduction
	IsIndependent bool      `json:"is_independent"`
	IsAuditable   bool      `json:"is_auditable"`
	IsSpecific    bool      `json:"is_specific"`
	CreatedAt     time.Time `json:"created_at"`
}
