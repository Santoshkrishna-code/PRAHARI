package consequence

import "time"

// Consequence describes potential outcome severity of a hazard scenario.
type Consequence struct {
	ID                 string    `json:"id"`
	ScenarioID         string    `json:"scenario_id"`
	Category           string    `json:"category"` // SAFETY, ENVIRONMENTAL, FINANCIAL, REPUTATIONAL, ASSET_DAMAGE
	Description        string    `json:"description"`
	SeverityScore      int       `json:"severity_score"`
	UnmitigatedRiskRank int      `json:"unmitigated_risk_rank"`
	CreatedAt          time.Time `json:"created_at"`
}
