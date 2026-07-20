package kafka

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
)

// GetDialer configures sasl dial options for the Kafka client.
func GetDialer(cfg Config) (*kafka.Dialer, error) {
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	if cfg.SASL.Enabled {
		var mechanism sasl.Mechanism
		var err error

		switch cfg.SASL.Mechanism {
		case "PLAIN":
			mechanism = plain.Mechanism{
				Username: cfg.SASL.Username,
				Password: cfg.SASL.Password,
			}
		case "SCRAM-SHA-256":
			mechanism, err = scram.Mechanism(
				scram.SHA256,
				cfg.SASL.Username,
				cfg.SASL.Password,
			)
			if err != nil {
				return nil, fmt.Errorf("failed to configure SCRAM-SHA-256: %w", err)
			}
		default:
			return nil, fmt.Errorf("unsupported SASL mechanism: %s", cfg.SASL.Mechanism)
		}

		dialer.SASLMechanism = mechanism
		dialer.TLS = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	return dialer, nil
}
