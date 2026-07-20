package optimization

import "time"

// Recommendation represents a water optimization suggestion.
type Recommendation struct {
	ID              string    `json:"id"`
	PlantID         string    `json:"plant_id"`
	AssetID         string    `json:"asset_id,omitempty"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	EstSavingKLD    float64   `json:"est_saving_kld"`
	EstSavingUSD    float64   `json:"est_saving_usd"`
	Priority        string    `json:"priority"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
}
