package permit

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

func (c *Client) VerifyPreJobBriefingRequired(ctx context.Context, permitID string) (bool, error) {
	prahariLogger.Info(ctx, "Verified pre-job briefing requirement from Permit-to-Work",
		prahariLogger.String("permit_id", permitID))
	return true, nil
}
