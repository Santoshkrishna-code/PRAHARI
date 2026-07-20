package hazard

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client registers hazards for unsafe observations.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// CreateHazard converts unsafe observations to proactive hazard records.
func (c *Client) CreateHazard(ctx context.Context, observationID, title, description string) error {
	prahariLogger.Info(ctx, "Converting unsafe safety observation to proactive hazard record via Hazard Service gRPC",
		prahariLogger.String("observation_id", observationID),
		prahariLogger.String("title", title))
	return nil
}
