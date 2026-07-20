package kafka

import (
	"context"
	"encoding/json"

	prahariLogger "prahari/shared/logger"
)

// InspectionCompletedEvent represents inbound inspection findings.
type InspectionCompletedEvent struct {
	InspectionID string   `json:"inspection_id"`
	Findings     []string `json:"findings"`
	AssetID      string   `json:"asset_id"`
}

// HazardStatusUpdater updates hazard status based on incoming messages.
type HazardStatusUpdater interface {
	TransitionStatus(ctx context.Context, cmd struct {
		HazardID   string
		TargetCode string
		ActorID    string
		Reason     string
	}) error
}

// Consumer handles inbound message consumer groups.
type Consumer struct {
	updater HazardStatusUpdater
}

// NewConsumer instantiates Consumer.
func NewConsumer(updater HazardStatusUpdater) *Consumer {
	return &Consumer{updater: updater}
}

// HandleInspectionCompleted processes findings.
func (c *Consumer) HandleInspectionCompleted(ctx context.Context, data []byte) error {
	var event InspectionCompletedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed inspection completed event, scanning findings for proactive hazard creation",
		prahariLogger.String("inspection_id", event.InspectionID),
		prahariLogger.Int("findings_count", len(event.Findings)))

	return nil
}
