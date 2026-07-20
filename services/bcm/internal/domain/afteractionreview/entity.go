package afteractionreview

import "time"

// Review represents an After Action Review conducted after BCP activation or exercise.
type Review struct {
	ID           string    `json:"id"`
	PlanID       string    `json:"plan_id"`
	ExerciseID   string    `json:"exercise_id,omitempty"`
	Title        string    `json:"title"`
	Summary      string    `json:"summary"`
	WhatWentWell string    `json:"what_went_well"`
	Improvements string    `json:"improvements"`
	Facilitator  string    `json:"facilitator_id"`
	ReviewedAt   time.Time `json:"reviewed_at"`
}
