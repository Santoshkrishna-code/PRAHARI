package retry

import "time"

// Message represents a failed transmission task scheduled for retry queues.
type Message struct {
	ID          string    `json:"id"`
	Payload     string    `json:"payload"`
	TargetTopic string    `json:"target_topic"`
	RetryCount  int       `json:"retry_count"`
	MaxRetries  int       `json:"max_retries"`
	NextRunAt   time.Time `json:"next_run_at"`
	CreatedAt   time.Time `json:"created_at"`
}
