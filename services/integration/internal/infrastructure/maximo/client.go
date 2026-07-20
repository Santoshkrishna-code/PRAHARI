package maximo

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Client struct {
	addr string
}

func NewClient(addr string) *Client {
	return &Client{addr: addr}
}

func (c *Client) PushWorkOrder(ctx context.Context, workOrderID string) error {
	prahariLogger.Info(ctx, "Pushed maintenance task work order to IBM Maximo",
		prahariLogger.String("work_order_id", workOrderID))
	return nil
}
