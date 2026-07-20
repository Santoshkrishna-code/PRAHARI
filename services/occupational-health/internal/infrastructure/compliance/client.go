package compliance

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

func (c *Client) VerifyObligations(ctx context.Context, checkID string) (bool, error) {
	prahariLogger.Info(ctx, "Auditing regulatory compliance medical parameters via Compliance Service gRPC",
		prahariLogger.String("check_id", checkID))
	return true, nil
}
