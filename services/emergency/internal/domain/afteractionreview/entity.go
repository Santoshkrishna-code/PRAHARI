package afteractionreview

import "time"

// Review represents an After Action Review (AAR) conducted following an emergency or drill.
type Review struct {
	ID          string    `json:"id"`
	EmergencyID string    `json:"emergency_id,omitempty"`
	DrillID     string    `json:"drill_id,omitempty"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	WhatWentWell string   `json:"what_went_well"`
	Improvements TEXT     `json:"improvements"`
	Facilitator string    `json:"facilitator_id"`
	ReviewedAt  time.Time `json:"reviewed_at"`
}

type TEXT = string
