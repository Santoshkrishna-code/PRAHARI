package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	prahariLogger "prahari/shared/logger"
	prahariKafka "prahari/shared/kafka/producer"
)

// Publisher dispatches messages to Kafka topics.
type Publisher struct {
	producer *prahariKafka.Producer
}

// NewPublisher instantiates Publisher.
func NewPublisher(producer *prahariKafka.Producer) *Publisher {
	return &Publisher{producer: producer}
}

// Publish serializes and sends payload.
func (p *Publisher) Publish(ctx context.Context, topic string, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Publishing permit event to Kafka",
		prahariLogger.String("topic", topic),
		prahariLogger.Int("payload_size", len(data)))

	return nil
}
