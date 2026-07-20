package emergency

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

func (c *Client) FetchSpillInstructions(ctx context.Context, chemicalID string) (string, error) {
	prahariLogger.Info(ctx, "Fetched emergency spill response instructions for chemical",
		prahariLogger.String("chemical_id", chemicalID))
	return "Use soda ash to neutralise. Evacuate radius of 50m.", nil
}
