package model

import "time"

// Model represents a registered computer vision model.
type Model struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"` // E.g., YOLOv8-PPE, SSD-Fire
	TaskType  string    `json:"task_type"` // DETECTION, SEGMENTATION, CLASSIFICATION
	Version   string    `json:"version"`
	UpdatedAt time.Time `json:"updated_at"`
}
