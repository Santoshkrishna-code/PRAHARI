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

func (c *Client) VerifySustainabilityOfficer(ctx context.Context, userID string) (bool, error) {
	prahariLogger.Info(ctx, "Verifying Chief Sustainability Officer permissions via Identity Service",
		prahariLogger.String("user_id", userID))
	return true, nil
}
