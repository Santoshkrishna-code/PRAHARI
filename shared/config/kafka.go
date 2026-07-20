package config

// KafkaConfig configures Apache Kafka / Amazon MSK connection parameters.
type KafkaConfig struct {
	Brokers          []string `env:"KAFKA_BROKERS" envSeparator:"," validate:"required,gt=0"`
	ClientID         string   `env:"KAFKA_CLIENT_ID" validate:"required"`
	GroupID          string   `env:"KAFKA_GROUP_ID"`
	SecurityProtocol string   `env:"KAFKA_SECURITY_PROTOCOL" envDefault:"PLAINTEXT" validate:"oneof=PLAINTEXT SSL SASL_PLAINTEXT SASL_SSL"`
	SASLMechanism    string   `env:"KAFKA_SASL_MECHANISM" validate:"oneof='' PLAIN SCRAM-SHA-256 SCRAM-SHA-512"`
	SASLUsername     string   `env:"KAFKA_SASL_USERNAME"`
	SASLPassword     string   `env:"KAFKA_SASL_PASSWORD"`
}
