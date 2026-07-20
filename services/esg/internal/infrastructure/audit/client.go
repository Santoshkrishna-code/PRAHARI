package audit

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

func (c *Client) FetchSustainabilityFindings(ctx context.Context, buID string) ([]string, error) {
	prahariLogger.Info(ctx, "Retrieving ESG governance finding metrics from Audit Service gRPC",
		prahariLogger.String("business_unit_id", buID))
	return []string{}, nil
}
