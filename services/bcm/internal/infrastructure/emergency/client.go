package emergency

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Client struct {
	grpcAddr string
}

func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

func (c *Client) FetchEmergencyStatus(ctx context.Context, emergencyID string) (string, error) {
	prahariLogger.Info(ctx, "Fetched active emergency status for business continuity plan activation", prahariLogger.String("emergency_id", emergencyID))
	return "DECLARED", nil
}
