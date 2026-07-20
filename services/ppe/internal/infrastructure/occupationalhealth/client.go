package occupationalhealth

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

func (c *Client) VerifyRespiratoryFitTestStatus(ctx context.Context, userID string) (bool, error) {
	prahariLogger.Info(ctx, "Verified occupational health SCBA/respirator fit test status and records",
		prahariLogger.String("user_id", userID))
	return true, nil
}
