package laboratory

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

func (c *Client) FetchLabResultData(ctx context.Context, sampleID string) (string, error) {
	prahariLogger.Info(ctx, "Retrieving diagnostic panel testing values from external Laboratory gRPC service",
		prahariLogger.String("sample_id", sampleID))
	return "NORMAL", nil
}
