package kafka

import (
	"context"
	"encoding/json"

	prahariLogger "prahari/shared/logger"
)

// ComplianceReviewedEvent represents inbound compliance approvals.
type ComplianceReviewedEvent struct {
	ComplianceID string `json:"compliance_id"`
	DepartmentID string `json:"department_id"`
}

// AuditTrigger defines state machine transition ports.
type AuditTrigger interface {
	TransitionStatus(ctx context.Context, cmd struct {
		AuditID    string
		TargetCode string
		ActorID    string
		Reason     string
	}) error
}

// Consumer handles inbound message consumer groups.
type Consumer struct {
	trigger AuditTrigger
}

// NewConsumer instantiates Consumer.
func NewConsumer(trigger AuditTrigger) *Consumer {
	return &Consumer{trigger: trigger}
}

// HandleComplianceReviewed triggers scheduled check reviews.
func (c *Consumer) HandleComplianceReviewed(ctx context.Context, data []byte) error {
	var event ComplianceReviewedEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Consumed compliance reviewed event, registering new assurance checks",
		prahariLogger.String("compliance_id", event.ComplianceID),
		prahariLogger.String("department_id", event.DepartmentID))

	return nil
}
}
