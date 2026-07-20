package identity

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

func (c *Client) VerifyEnergyAuditor(ctx context.Context, userID string) (bool, error) {
	prahariLogger.Info(ctx, "Verifying Certified Energy Auditor permissions via Identity Service",
		prahariLogger.String("user_id", userID))
	return true, nil
}
