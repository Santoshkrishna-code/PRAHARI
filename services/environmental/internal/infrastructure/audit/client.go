package audit

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

func (c *Client) FetchAuditFindings(ctx context.Context, siteID string) ([]string, error) {
	prahariLogger.Info(ctx, "Retrieving site ISO 14001 audit findings list from Audit Service gRPC",
		prahariLogger.String("site_id", siteID))
	return []string{}, nil
}
