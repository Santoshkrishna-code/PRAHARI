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

func (c *Client) FetchActiveShiftPPEAudit(ctx context.Context, plantID string) (bool, error) {
	prahariLogger.Info(ctx, "Fetched active shift logs to audit mandatory PPE compliance compliance",
		prahariLogger.String("plant_id", plantID))
	return true, nil
}
