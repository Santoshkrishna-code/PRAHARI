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

func (c *Client) FetchChemicalAnalysisData(ctx context.Context, sampleID string) (float64, error) {
	prahariLogger.Info(ctx, "Retrieving chemical test panel analysis results via Laboratory Service gRPC",
		prahariLogger.String("sample_id", sampleID))
	return 0.05, nil
}
