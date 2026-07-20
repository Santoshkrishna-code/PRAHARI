package kafka

import (
	"context"
	"encoding/json"

	prahariLogger "prahari/shared/logger"
)

// InspectionEvent represents incoming message updates schemas.
type InspectionEvent struct {
	InspectionID string `json:"inspection_id"`
	AssetID      string `json:"asset_id"`
	Status       string `json:"status"`
}

// AssetStatusUpdater updates asset lifecycle based on inbound messages.
type AssetStatusUpdater interface {
	TransitionLifecycle(ctx context.Context, cmd struct {
		AssetID    string
		TargetCode string
		ActorID    string
		Reason     string
	}) error
}

// Consumer handles inbound message consumer groups.
type Consumer struct {
	updater AssetStatusUpdater
}

// NewConsumer instantiates Consumer.
func NewConsumer(updater AssetStatusUpdater) *Consumer {
	return &Consumer{updater: updater}
}

// HandleInspectionCompleted processes completion checks.
func (c *Consumer) HandleInspectionCompleted(ctx context.Context, data []byte) error {
	var event InspectionEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed inspection completion event for asset",
		prahariLogger.String("inspection_id", event.InspectionID),
		prahariLogger.String("asset_id", event.AssetID))

	return nil
}
