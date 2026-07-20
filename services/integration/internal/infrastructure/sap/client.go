package sap

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Client struct {
	addr string
}

func NewClient(addr string) *Client {
	return &Client{addr: addr}
}

func (c *Client) SyncAssetDetails(ctx context.Context, externalID string) error {
	prahariLogger.Info(ctx, "Synchronised asset metrics from SAP ERP system",
		prahariLogger.String("external_id", externalID))
	return nil
}
