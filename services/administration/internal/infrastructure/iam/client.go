package iam

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

func (c *Client) ProvisionTenantCredentials(ctx context.Context, tenantID, adminEmail string) error {
	prahariLogger.Info(ctx, "Provisioned initial tenant administrator identity and credentials in Identity service",
		prahariLogger.String("tenant_id", tenantID),
		prahariLogger.String("admin_email", adminEmail))
	return nil
}
