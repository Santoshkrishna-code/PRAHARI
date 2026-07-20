package workflow

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

func (c *Client) TriggerTenantOnboarding(ctx context.Context, tenantID string) error {
	prahariLogger.Info(ctx, "Triggered tenant provisioning and onboarding workflow in Workflow Engine",
		prahariLogger.String("tenant_id", tenantID))
	return nil
}
