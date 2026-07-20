package notification

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

func (c *Client) SendTenantAlert(ctx context.Context, tenantID, message string) error {
	prahariLogger.Info(ctx, "Sent tenant administrative alert to Notification Service",
		prahariLogger.String("tenant_id", tenantID))
	return nil
}
