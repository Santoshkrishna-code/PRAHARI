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

func (c *Client) TriggerDocumentApprovalWorkflow(ctx context.Context, docID string) error {
	prahariLogger.Info(ctx, "Triggered document multi-level approval workflow in Workflow Engine", prahariLogger.String("document_id", docID))
	return nil
}
