package document

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

func (c *Client) TriggerPAndIDRevision(ctx context.Context, docNumber string) error {
	prahariLogger.Info(ctx, "Triggered document revision workflow for P&ID / SOP", prahariLogger.String("document_number", docNumber))
	return nil
}
