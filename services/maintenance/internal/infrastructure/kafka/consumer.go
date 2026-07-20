package kafka

import (
	"context"
	"encoding/json"

	prahariLogger "prahari/shared/logger"
)

// PermitApprovedEvent represents inbound work permit approvals.
type PermitApprovedEvent struct {
	PermitID string `json:"permit_id"`
	AssetID  string `json:"asset_id"`
	Status   string `json:"status"`
}

// MaintenanceStatusUpdater updates maintenance status based on incoming messages.
type MaintenanceStatusUpdater interface {
	TransitionStatus(ctx context.Context, cmd struct {
		MaintenanceID string
		TargetCode    string
		ActorID       string
		Reason        string
	}) error
}

// Consumer handles inbound message consumer groups.
type Consumer struct {
	updater MaintenanceStatusUpdater
}

// NewConsumer instantiates Consumer.
func NewConsumer(updater MaintenanceStatusUpdater) *Consumer {
	return &Consumer{updater: updater}
}

// HandlePermitApproved processes permit confirmations.
func (c *Consumer) HandlePermitApproved(ctx context.Context, data []byte) error {
	var event PermitApprovedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed permit approval event for asset",
		prahariLogger.String("permit_id", event.PermitID),
		prahariLogger.String("asset_id", event.AssetID))

	return nil
}
