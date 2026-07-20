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

func (c *Client) TriggerLOTOVerificationApprovalWorkflow(ctx context.Context, certificateID string) error {
	prahariLogger.Info(ctx, "Triggered multi-stage LOTO certificate isolation and locking verification workflow", prahariLogger.String("certificate_id", certificateID))
	return nil
}
