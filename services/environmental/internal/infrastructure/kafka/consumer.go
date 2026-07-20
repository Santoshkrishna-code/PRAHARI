package kafka

import (
	"context"
	"encoding/json"

	prahariLogger "prahari/shared/logger"
)

type IncidentCreatedEvent struct {
	IncidentID   string `json:"incident_id"`
	PlantID      string `json:"plant_id"`
	IncidentType string `json:"incident_type"` // "SPILL", "GAS_RELEASE"
}

type MaintenanceCompletedEvent struct {
	AssetID      string `json:"asset_id"`
	WorkOrderID  string `json:"work_order_id"`
	CompletionAt string `json:"completion_at"`
}

type RiskActivatedEvent struct {
	RiskID        string `json:"risk_id"`
	HazardCode    string `json:"hazard_code"`
	SeverityScale int    `json:"severity_scale"`
}

type Consumer struct{}

func NewConsumer() *Consumer {
	return &Consumer{}
}

func (c *Consumer) HandleIncidentCreated(ctx context.Context, data []byte) error {
	var event IncidentCreatedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed incident created event, triggering environmental risk assessments evaluations",
		prahariLogger.String("incident_id", event.IncidentID),
		prahariLogger.String("type", event.IncidentType))

	return nil
}

func (c *Consumer) HandleMaintenanceCompleted(ctx context.Context, data []byte) error {
	var event MaintenanceCompletedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed maintenance completed event, checking parameters on emission sources",
		prahariLogger.String("asset_id", event.AssetID))

	return nil
}

func (c *Consumer) HandleRiskActivated(ctx context.Context, data []byte) error {
	var event RiskActivatedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed risk activated event, raising frequency parameters on sampling runs",
		prahariLogger.String("risk_id", event.RiskID))

	return nil
}
