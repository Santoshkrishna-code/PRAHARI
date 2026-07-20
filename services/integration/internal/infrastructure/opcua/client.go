package opcua

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Client struct {
	endpoint string
}

func NewClient(endpoint string) *Client {
	return &Client{endpoint: endpoint}
}

func (c *Client) ReadNodeValue(ctx context.Context, nodeID string) (float64, error) {
	prahariLogger.Info(ctx, "Read industrial sensor metrics via OPC UA protocol node ID",
		prahariLogger.String("node_id", nodeID))
	return 101.5, nil
}
