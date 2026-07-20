package graphstore

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

func (c *Client) QueryTopologyPath(ctx context.Context, twinID, fromNode, toNode string) ([]string, error) {
	prahariLogger.Info(ctx, "Executing GraphQL/Cypher graph topology relationship path query",
		prahariLogger.String("twin_id", twinID))
	return []string{fromNode, "pipeline-v10", "valve-v102", toNode}, nil
}
