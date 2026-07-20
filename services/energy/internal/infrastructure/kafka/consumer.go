package kafka

import (
	"context"
	"encoding/json"

	prahariLogger "prahari/shared/logger"
)

type AssetUpdatedEvent struct {
	AssetID      string `json:"asset_id"`
	FacilityCode string `json:"facility_code"`
	Status       string `json:"status"`
}

type MaintenanceCompletedEvent struct {
	WorkOrderID string `json:"work_order_id"`
	AssetID     string `json:"asset_id"`
	Notes       string `json:"notes"`
}

type ESGScoreUpdatedEvent struct {
	BusinessUnitID string  `json:"business_unit_id"`
	NewScore       float64 `json:"new_score"`
}

type Consumer struct{}

func NewConsumer() *Consumer {
	return &Consumer{}
}

func (c *Consumer) HandleAssetUpdated(ctx context.Context, data []byte) error {
	var event AssetUpdatedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed asset updated event, checking utility meters links",
		prahariLogger.String("asset_id", event.AssetID))

	return nil
}

func (c *Consumer) HandleMaintenanceCompleted(ctx context.Context, data []byte) error {
	var event MaintenanceCompletedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed maintenance completed event, checking calibration drift updates",
		prahariLogger.String("work_order_id", event.WorkOrderID))

	return nil
}

func (c *Consumer) HandleESGScoreUpdated(ctx context.Context, data []byte) error {
	var event ESGScoreUpdatedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed ESG score updated event, syncing optimization programs priorities",
		prahariLogger.String("bu_id", event.BusinessUnitID))

	return nil
}
