package handler

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Request defines the input event payload structure.
type Request struct {
	EventID string `json:"event_id"`
}

// HandleRequest processes the incoming serverless invocation.
func HandleRequest(ctx context.Context, req Request) (string, error) {
	prahariLogger.Info(ctx, "AWS Lambda event execution processed", prahariLogger.String("event_id", req.EventID))
	return "execution-success", nil
}
