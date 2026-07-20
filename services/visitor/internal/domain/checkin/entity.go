package checkin

import "time"

// Record tracks physical plant check-in event details.
type Record struct {
	ID                 string    `json:"id"`
	VisitID            string    `json:"visit_id"`
	SecurityCheckPoint string    `json:"security_check_point"`
	GateNumber         string    `json:"gate_number"`
	CheckInAt          time.Time `json:"check_in_at"`
	CheckedInBy        string    `json:"checked_in_by"`
}
