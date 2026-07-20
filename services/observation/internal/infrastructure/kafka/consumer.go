package kafka

import (
	"context"
	"encoding/json"

	prahariLogger "prahari/shared/logger"
)

// NearMissCreatedEvent represents inbound near miss logs.
type NearMissCreatedEvent struct {
	NearMissID string `json:"near_miss_id"`
	Title      string `json:"title"`
	AssetID    string `json:"asset_id"`
}

// ObservationStatusUpdater updates status based on incoming messages.
type ObservationStatusUpdater interface {
	TransitionStatus(ctx context.Context, cmd struct {
		ObservationID string
		TargetCode    string
		ActorID       string
		Reason        string
	}) error
}

// Consumer handles inbound message consumer groups.
type Consumer struct {
	updater ObservationStatusUpdater
}

// NewConsumer instantiates Consumer.
func NewConsumer(updater ObservationStatusUpdater) *Consumer {
	return &Consumer{updater: updater}
}

// HandleNearMissCreated reviews near miss creations.
func (c *Consumer) HandleNearMissCreated(ctx context.Context, data []byte) error {
	var event NearMissCreatedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed near miss created event, scanning behavioral trends for auto-coaching requirements",
		prahariLogger.String("near_miss_id", event.NearMissID),
		prahariLogger.String("asset_id", event.AssetID))

	return nil
}
