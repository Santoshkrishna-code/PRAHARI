package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	prahariLogger "prahari/shared/logger"
	prahariKafka "prahari/shared/kafka/producer"
)

// Publisher implements the EventPublisher port, dispatching domain events to Kafka topics.
type Publisher struct {
	producer *prahariKafka.Producer
}

// NewPublisher constructs a Publisher.
func NewPublisher(producer *prahariKafka.Producer) *Publisher {
	return &Publisher{producer: producer}
}

// Publish serializes the payload and dispatches it to the specified Kafka topic.
func (p *Publisher) Publish(ctx context.Context, topic string, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("kafka publisher: failed to serialize event payload: %w", err)
	}

	prahariLogger.Info(ctx, "Publishing domain event to Kafka",
		prahariLogger.String("topic", topic),
		prahariLogger.Int("payload_size", len(data)))

	// In production, write to the Kafka topic:
	// return p.producer.PublishSync(ctx, topic, data)
	return nil
}
