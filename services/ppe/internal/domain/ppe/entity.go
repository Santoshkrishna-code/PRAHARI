package ppe

import "time"

// PPE represents a type of Personal Protective Equipment in the plant catalog.
type PPE struct {
	ID          string    `json:"id"`
	PlantID     string    `json:"plant_id"`
	ModelName   string    `json:"model_name"` // E.g. Standard Hard Hat, Arc Flash Suit Class 4
	CategoryID  string    `json:"category_id"`
	Manufacturer string   `json:"manufacturer"`
	PartNumber  string    `json:"part_number"`
	StandardRef string    `json:"standard_ref"` // E.g. ANSI Z89.1, EN 397
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
