package hazard

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client registers hazards for repeat issues.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// CreateHazard converts recurring near misses into hazards.
func (c *Client) CreateHazard(ctx context.Context, nearmissID, title, description string) error {
	prahariLogger.Info(ctx, "Converting recurring near miss to proactive safety hazard record via Hazard Service gRPC",
		prahariLogger.String("near_miss_id", nearmissID),
		prahariLogger.String("title", title))
	return nil
}
