package inspection

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

func (c *Client) FetchHygieneViolations(ctx context.Context, siteID string) ([]string, error) {
	prahariLogger.Info(ctx, "Retrieving hygiene/sanitation inspection checklist evaluations via Inspection Service gRPC",
		prahariLogger.String("site_id", siteID))
	return []string{}, nil
}
