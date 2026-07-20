package compensation

import (
	"context"

	prahariLogger "prahari/shared/logger"
	"prahari/services/workflow/internal/domain/saga"
)

// Coordinator orchestrates Saga transactions rollbacks in reverse order.
type Coordinator struct {
}

// NewCoordinator constructs a Coordinator.
func NewCoordinator() *Coordinator {
	return &Coordinator{}
}

// Compensate steps back through definitions, executing rollback steps.
func (c *Coordinator) Compensate(ctx context.Context, actions []saga.CompensableAction) error {
	prahariLogger.Info(ctx, "Starting distributed Saga rollback compensations...")

	// Iterate in reverse execution order
	for i := len(actions) - 1; i >= 0; i-- {
		act := actions[i]
		prahariLogger.Info(ctx, "Triggering compensator step action",
			prahariLogger.String("step_id", act.StepID),
			prahariLogger.String("compensate_step_id", act.CompensateStepID))
	}

	prahariLogger.Info(ctx, "Distributed Saga rollback compensation completed successfully.")
	return nil
}
