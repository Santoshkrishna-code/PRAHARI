package processnode

import "time"

// Node represents a defined section of a process design isolatable for HAZOP analysis.
type Node struct {
	ID           string    `json:"id"`
	StudyID      string    `json:"study_id"`
	NodeNumber   int       `json:"node_number"`
	NodeName     string    `json:"node_name"`
	DesignIntent string    `json:"design_intent"`
	PAndIDNumber string    `json:"p_and_id_number"`
	OperatingTemp float64  `json:"operating_temp_c"`
	OperatingPress float64 `json:"operating_press_bar"`
	LocationCode string    `json:"location_code"`
	CreatedAt    time.Time `json:"created_at"`
}
