package kafka

import (
	"context"
	"encoding/json"

	prahariLogger "prahari/shared/logger"
)

// PermitCreatedEvent represents inbound work permit activations.
type PermitCreatedEvent struct {
	PermitID     string `json:"permit_id"`
	ContractorID string `json:"contractor_id"`
	Status       string `json:"status"`
}

// ContractorStatusUpdater updates contractor status based on incoming messages.
type ContractorStatusUpdater interface {
	TransitionStatus(ctx context.Context, cmd struct {
		ContractorID string
		TargetCode   string
		ActorID      string
		Reason       string
	}) error
}

// Consumer handles inbound message consumer groups.
type Consumer struct {
	updater ContractorStatusUpdater
}

// NewConsumer instantiates Consumer.
func NewConsumer(updater ContractorStatusUpdater) *Consumer {
	return &Consumer{updater: updater}
}

// HandlePermitCreated processes permit registration event.
func (c *Consumer) HandlePermitCreated(ctx context.Context, data []byte) error {
	var event PermitCreatedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed permit creation event for contractor company",
		prahariLogger.String("permit_id", event.PermitID),
		prahariLogger.String("contractor_id", event.ContractorID))

	return nil
}
