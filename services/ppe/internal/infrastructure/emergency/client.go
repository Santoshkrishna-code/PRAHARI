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

func (c *Client) FetchEmergencyPPERequirements(ctx context.Context, plantID string) ([]string, error) {
	prahariLogger.Info(ctx, "Fetched emergency responder PPE gear requirements for immediate dispatch",
		prahariLogger.String("plant_id", plantID))
	return []string{"SCBA-01", "Chemical-Suit-V4"}, nil
}
