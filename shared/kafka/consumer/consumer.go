package consumer

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	prahariKafka "prahari/shared/kafka"
)

// Consumer wraps segmentio/kafka-go Reader to read topic partitions.
type Consumer struct {
	reader *kafka.Reader
}

// NewConsumer constructs a consumer reading from a specific topic.
func NewConsumer(cfg prahariKafka.Config, topic string) (*Consumer, error) {
	dialer, err := prahariKafka.GetDialer(cfg)
	if err != nil {
		return nil, err
	}

	transport := &kafka.Transport{
		Dial:   dialer.DialFunc,
		SASL:   dialer.SASLMechanism,
		TLS:    dialer.TLS,
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   cfg.Brokers,
		GroupID:   cfg.GroupID,
		Topic:     topic,
		Dialer:    dialer,
		Transport: transport,
		MaxBytes:  10e6, // 10MB limit
	})

	return &Consumer{
		reader: reader,
	}, nil
}

// StartConsumer runs a blocking read-commit loop.
// Message offsets are committed only if the handler completes without error (at-least-once guarantee).
func (c *Consumer) StartConsumer(ctx context.Context, handler func(ctx context.Context, key, val []byte) error) error {
	if c.reader == nil {
		return fmt.Errorf("consumer reader is uninitialized")
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// Fetch message blocks (does not auto-commit offsets)
			msg, err := c.reader.FetchMessage(ctx)
			if err != nil {
				// Return errors to trigger reconnect loops
				return fmt.Errorf("failed to fetch message from Kafka: %w", err)
			}

			// Process message
			err = handler(ctx, msg.Key, msg.Value)
			if err == nil {
				// Commit offset on successful processing
				err = c.reader.CommitMessages(ctx, msg)
				if err != nil {
					return fmt.Errorf("failed to commit offset for message: %w", err)
				}
			}
		}
	}
}

// Close closes reader connections.
func (c *Consumer) Close() error {
	if c.reader == nil {
		return nil
	}
	return c.reader.Close()
}
