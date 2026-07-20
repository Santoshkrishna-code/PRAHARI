package exercise

import "time"

// Exercise represents a tabletop or full-scale emergency preparedness exercise.
type Exercise struct {
	ID           string    `json:"id"`
	PlantID      string    `json:"plant_id"`
	Title        string    `json:"title"`
	ExerciseType string    `json:"exercise_type"` // TABLETOP, FUNCTIONAL, FULL_SCALE
	Objectives   string    `json:"objectives"`
	ConductedAt  time.Time `json:"conducted_at"`
	CreatedAt    time.Time `json:"created_at"`
}
