package kafka

import (
	"context"
	"encoding/json"

	prahariLogger "prahari/shared/logger"
)

type EnvironmentClosedEvent struct {
	AspectID   string `json:"aspect_id"`
	PlantID    string `json:"plant_id"`
	Resolution string `json:"resolution"`
}

type ComplianceFailedEvent struct {
	SourceType string `json:"source_type"`
	Reason     string `json:"reason"`
	Severity   string `json:"severity"`
}

type AuditClosedEvent struct {
	AuditID      string   `json:"audit_id"`
	Findings     []string `json:"findings"`
	CompletionAt string   `json:"completion_at"`
}

type Consumer struct{}

func NewConsumer() *Consumer {
	return &Consumer{}
}

func (c *Consumer) HandleEnvironmentClosed(ctx context.Context, data []byte) error {
	var event EnvironmentClosedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed environment closed event, updating sustainability indicators",
		prahariLogger.String("aspect_id", event.AspectID))

	return nil
}

func (c *Consumer) HandleComplianceFailed(ctx context.Context, data []byte) error {
	var event ComplianceFailedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed compliance failed event, recalculating ESG scorecards factors",
		prahariLogger.String("source", event.SourceType))

	return nil
}

func (c *Consumer) HandleAuditClosed(ctx context.Context, data []byte) error {
	var event AuditClosedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed audit closed event, verifying target initiatives",
		prahariLogger.String("audit_id", event.AuditID))

	return nil
}
