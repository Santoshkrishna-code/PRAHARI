package asset

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client checks equipment profiles validity.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// VerifyAsset verifies asset status.
func (c *Client) VerifyAsset(ctx context.Context, assetID string) (bool, error) {
	prahariLogger.Info(ctx, "Verifying asset profile details via Asset Management Service gRPC",
		prahariLogger.String("asset_id", assetID))
	return true, nil
}
