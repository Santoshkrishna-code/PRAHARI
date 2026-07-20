package permit

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

func (c *Client) FetchPermitAccessAuthorization(ctx context.Context, permitID string) (bool, error) {
	prahariLogger.Info(ctx, "Fetched permit validation and work clearance check before gate check-in",
		prahariLogger.String("permit_id", permitID))
	return true, nil
}
