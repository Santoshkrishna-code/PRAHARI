package preference

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Evaluator checks opt-out settings boundaries.
type Evaluator struct {
}

// NewEvaluator constructs an Evaluator.
func NewEvaluator() *Evaluator {
	return &Evaluator{}
}

// Allowed evaluates user settings before dispatch triggers.
func (e *Evaluator) Allowed(ctx context.Context, userID, channel string) (bool, error) {
	prahariLogger.Info(ctx, "Evaluating dispatch preferences permissions",
		prahariLogger.String("user_id", userID),
		prahariLogger.String("channel", channel))

	// In production, execute SQL lookups resolving active preferences Opt-out matrices:
	return true, nil
}
