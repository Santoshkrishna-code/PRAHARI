package events

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.uber.org/zap"
)

// EventPublisher defines the port for publishing events to a message broker.
type EventPublisher interface {
	Publish(ctx context.Context, topic string, payload interface{}) error
}

// LogEventPublisher implements EventPublisher by dumping events to structured logs.
// Can be replaced by KafkaEventPublisher.
type LogEventPublisher struct {
	logger *zap.Logger
}

func NewLogEventPublisher(logger *zap.Logger) EventPublisher {
	return &LogEventPublisher{logger: logger}
}

func (p *LogEventPublisher) Publish(ctx context.Context, topic string, payload interface{}) error {
	bytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal event payload: %w", err)
	}

	p.logger.Info("Event Published",
		zap.String("topic", topic),
		zap.ByteString("payload", bytes),
		zap.Time("published_at", time.Now()),
	)
	return nil
}
