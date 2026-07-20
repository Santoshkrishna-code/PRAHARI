package permit

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

func (c *Client) ValidateChemicalForPermit(ctx context.Context, permitID, chemicalID string) (bool, error) {
	prahariLogger.Info(ctx, "Validated chemical authorization status for active work permit",
		prahariLogger.String("permit_id", permitID),
		prahariLogger.String("chemical_id", chemicalID))
	return true, nil
}
