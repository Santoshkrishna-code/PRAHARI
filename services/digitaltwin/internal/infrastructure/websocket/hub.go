package websocket

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Hub struct{}

func NewHub() *Hub {
	return &Hub{}
}

func (h *Hub) BroadcastTwinState(ctx context.Context, twinID string, statePayload string) {
	prahariLogger.Info(ctx, "Broadcasting real-time Digital Twin state snapshot to active visualization clients",
		prahariLogger.String("twin_id", twinID))
}
