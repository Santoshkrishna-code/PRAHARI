package approval

import (
	"context"
	"errors"

	prahariLogger "prahari/shared/logger"
)

// Service coordinates dynamic decision approvals and updates step statuses.
type Service struct {
}

// NewService constructs an Approval Service.
func NewService() *Service {
	return &Service{}
}

// SubmitDecision records choices (approve/reject), evaluating quorum steps.
func (s *Service) SubmitDecision(ctx context.Context, taskID, approverID, decision string) error {
	if taskID == "" || approverID == "" || decision == "" {
		return errors.New("approval payload requires task ID, approver ID and decision fields")
	}

	prahariLogger.Info(ctx, "Approval decision received",
		prahariLogger.String("task_id", taskID),
		prahariLogger.String("approver_id", approverID),
		prahariLogger.String("decision", decision))

	return nil
}
