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

// PublishWorkflowStarted dispatches state events.
func (p *Publisher) PublishWorkflowStarted(ctx context.Context, instID string) error {
	// In production, build Envelope models and write to topic keys:
	// return p.prod.PublishSync(ctx, "workflow.started", []byte(payload))
	return nil
}
