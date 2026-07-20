package kafka

// SASLConfig models simple authentication properties for Kafka brokers.
type SASLConfig struct {
	Enabled   bool   `json:"enabled"`
	Mechanism string `json:"mechanism"` // PLAIN, SCRAM-SHA-256, etc.
	Username  string `json:"username"`
	Password  string `json:"password"`
}

// Config maps standard broker cluster details.
type Config struct {
	Brokers  []string   `json:"brokers"`
	ClientID string     `json:"client_id"`
	GroupID  string     `json:"group_id"`
	SASL     SASLConfig `json:"sasl"`
}
