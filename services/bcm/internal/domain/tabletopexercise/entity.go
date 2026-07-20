package tabletopexercise

import "time"

// Exercise tracks scenario-based walkthrough exercises with crisis management teams.
type Exercise struct {
	ID          string    `json:"id"`
	ExerciseID  string    `json:"exercise_id"`
	Scenario    string    `json:"scenario"`
	ModeratorID string    `json:"moderator_id"`
	ConductedAt time.Time `json:"conducted_at"`
	CreatedAt   time.Time `json:"created_at"`
}
