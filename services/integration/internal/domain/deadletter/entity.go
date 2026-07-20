package deadletter

import "time"

// Message holds payloads rejected permanently after exhausting all retries.
type Message struct {
	ID         string    `json:"id"`
	Payload    string    `json:"payload"`
	TopicName  string    `json:"topic_name"`
	ErrorMsg   string    `json:"error_msg"`
	ArchivedAt time.Time `json:"archived_at"`
}
