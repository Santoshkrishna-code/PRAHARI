package hazardscenario

import "time"

// Scenario represents a specific hazardous cause-consequence sequence.
type Scenario struct {
	ID             string    `json:"id"`
	NodeID         string    `json:"node_id"`
	DeviationID    string    `json:"deviation_id"`
	CauseDescription string  `json:"cause_description"`
	Severity       int       `json:"severity"`       // 1 (Low) to 5 (Catastrophic)
	Likelihood     int       `json:"likelihood"`     // 1 (Improbable) to 5 (Frequent)
	RiskRank       int       `json:"risk_rank"`      // Matrix product (1 to 25)
	RiskCategory   string    `json:"risk_category"`  // LOW, MEDIUM, HIGH, UNACCEPTABLE
	IsSILRelevant  bool      `json:"is_sil_relevant"`
	CreatedAt      time.Time `json:"created_at"`
}
