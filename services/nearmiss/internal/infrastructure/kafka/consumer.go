package kafka

import (
	"context"
	"encoding/json"

	prahariLogger "prahari/shared/logger"
)

// IncidentCreatedEvent represents inbound actual incident logs.
type IncidentCreatedEvent struct {
	IncidentID string `json:"incident_id"`
	Title      string `json:"title"`
	AssetID    string `json:"asset_id"`
}

// NearMissStatusUpdater updates status based on incoming messages.
type NearMissStatusUpdater interface {
	TransitionStatus(ctx context.Context, cmd struct {
		NearMissID string
		TargetCode string
		ActorID    string
		Reason     string
	}) error
}

// Consumer handles inbound message consumer groups.
type Consumer struct {
	updater NearMissStatusUpdater
}

// NewConsumer instantiates Consumer.
func NewConsumer(updater NearMissStatusUpdater) *Consumer {
	return &Consumer{updater: updater}
}

// HandleIncidentCreated reviews incident creations.
func (c *Consumer) HandleIncidentCreated(ctx context.Context, data []byte) error {
	var event IncidentCreatedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed incident created event, reviewing near miss trends for auto-mitigation correlations",
		prahariLogger.String("incident_id", event.IncidentID),
		prahariLogger.String("asset_id", event.AssetID))

	return nil
}
