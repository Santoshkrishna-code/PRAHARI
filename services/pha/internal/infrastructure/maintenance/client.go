package maintenance

import (
	"context"
	"fmt"
	"time"

	prahariLogger "prahari/shared/logger"
)

type Client struct {
	grpcAddr string
}

func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

func (c *Client) CreateWorkOrder(ctx context.Context, plantID, title, description string) (string, error) {
	woID := fmt.Sprintf("wo-pha-%d", time.Now().UnixNano())
	prahariLogger.Info(ctx, "Created maintenance work order for PHA recommendation action item",
		prahariLogger.String("work_order_id", woID),
		prahariLogger.String("title", title))
	return woID, nil
}
