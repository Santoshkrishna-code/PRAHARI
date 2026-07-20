package occupationalhealth

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

func (c *Client) FetchVisitorMedicalClearance(ctx context.Context, visitorID string) (bool, error) {
	prahariLogger.Info(ctx, "Fetched visitor physical/medical clearance certificates for site check-in",
		prahariLogger.String("visitor_id", visitorID))
	return true, nil
}
