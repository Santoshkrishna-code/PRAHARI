package kafka

import (
	"context"

	prahariKafka "prahari/shared/kafka/producer"
)

// Publisher adapter dispatching platform event envelopes.
type Publisher struct {
	prod *prahariKafka.Producer
}

// NewPublisher constructs a Publisher.
func NewPublisher(prod *prahariKafka.Producer) *Publisher {
	return &Publisher{prod: prod}
}

// PublishUserCreated dispatches user signup signals.
func (p *Publisher) PublishUserCreated(ctx context.Context, userID, email string) error {
	// In production, build Envelope models and write to topic keys:
	// return p.prod.PublishSync(ctx, "identity.user.created", []byte(payload))
	return nil
}
