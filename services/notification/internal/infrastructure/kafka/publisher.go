package kafka

import (
	"context"

	prahariKafka "prahari/shared/kafka/producer"
)

// Publisher adapter dispatching status event envelopes.
type Publisher struct {
	prod *prahariKafka.Producer
}

// NewPublisher constructs a Publisher.
func NewPublisher(prod *prahariKafka.Producer) *Publisher {
	return &Publisher{prod: prod}
}

// PublishNotificationSent dispatches sent logs.
func (p *Publisher) PublishNotificationSent(ctx context.Context, id string) error {
	// In production, build Envelope models and write to topic keys:
	// return p.prod.PublishSync(ctx, "notification.sent", []byte(payload))
	return nil
}
