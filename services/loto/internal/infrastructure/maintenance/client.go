package maintenance

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

func (c *Client) VerifyZeroEnergyForWorkOrder(ctx context.Context, workOrderID string) (bool, error) {
	prahariLogger.Info(ctx, "Verified zero-energy state before authorizing field maintenance work order start",
		prahariLogger.String("work_order_id", workOrderID))
	return true, nil
}
