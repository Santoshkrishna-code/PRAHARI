package conversation

import "time"

// Thread represents a chat session log thread.
type Thread struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	PlantID   string    `json:"plant_id"`
	CreatedAt time.Time `json:"created_at"`
}

// Message represents a user or AI chat message payload.
type Message struct {
	ID        string    `json:"id"`
	ThreadID  string    `json:"thread_id"`
	Role      string    `json:"role"` // USER, ASSISTANT, SYSTEM
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
