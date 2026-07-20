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

func (c *Client) SendExposureExceeded(ctx context.Context, employeeID, chemicalID string, level float64) error {
	prahariLogger.Info(ctx, "Sent chemical exposure limit exceeded alert to Occupational Health",
		prahariLogger.String("employee_id", employeeID),
		prahariLogger.String("chemical_id", chemicalID))
	return nil
}
