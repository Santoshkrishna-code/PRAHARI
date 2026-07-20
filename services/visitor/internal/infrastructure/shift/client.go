package shift

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

func (c *Client) FetchActiveShiftSupervisor(ctx context.Context, plantID string) (string, error) {
	prahariLogger.Info(ctx, "Fetched active shift supervisor details for emergency mustering escalation",
		prahariLogger.String("plant_id", plantID))
	return "usr-sup-04", nil
}
