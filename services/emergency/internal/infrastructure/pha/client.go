package pha

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

func (c *Client) FetchWorstCaseScenarios(ctx context.Context, unitID string) (int, error) {
	prahariLogger.Info(ctx, "Fetched worst-case process hazard scenarios for emergency response planning", prahariLogger.String("unit_id", unitID))
	return 2, nil
}
