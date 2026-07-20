package execution

import (
	"context"
	"fmt"

	prahariLogger "prahari/shared/logger"
	"prahari/services/workflow/internal/domain/instance"
	"prahari/services/workflow/internal/domain/state"
	"prahari/services/workflow/internal/domain/workflow"
)

// Interpreter parses DSL definition steps, executing transition loops.
type Interpreter struct {
}

// NewInterpreter constructs an Interpreter.
func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

// Execute steps through transition runs sequentially.
func (i *Interpreter) Execute(ctx context.Context, inst *instance.Instance, def *workflow.Definition) error {
	inst.State = state.StateRunning
	prahariLogger.Info(ctx, "Starting workflow instance execution loop", prahariLogger.String("instance_id", inst.ID))

	stepMap := make(map[string]*workflow.Step)
	for idx := range def.Steps {
		stepMap[def.Steps[idx].ID] = &def.Steps[idx]
	}

	currID := inst.CurrentStepID
	if currID == "" && len(def.Steps) > 0 {
		currID = def.Steps[0].ID
	}

	for currID != "" {
		step, exists := stepMap[currID]
		if !exists {
			inst.State = state.StateFailed
			return fmt.Errorf("interpreter error: missing step ID %s", currID)
		}

		prahariLogger.Info(ctx, "Executing workflow step action",
			prahariLogger.String("step_id", step.ID),
			prahariLogger.String("type", step.Type))

		// In production, execute automated task actions here (email, script, etc.)
		currID = step.NextStep
		inst.CurrentStepID = currID
	}

	inst.State = state.StateCompleted
	prahariLogger.Info(ctx, "Workflow instance execution completed successfully", prahariLogger.String("instance_id", inst.ID))
	return nil
}
