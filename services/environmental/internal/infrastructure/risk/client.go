package risk

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

func (c *Client) TriggerEnvironmentalRiskAssessment(ctx context.Context, aspectID string) error {
	prahariLogger.Info(ctx, "Triggering environmental aspect risk evaluation with Risk Assessment Management Service",
		prahariLogger.String("aspect_id", aspectID))
	return nil
}
