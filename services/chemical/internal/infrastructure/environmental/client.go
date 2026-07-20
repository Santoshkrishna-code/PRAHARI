package environmental

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

func (c *Client) ReportChemicalDisposal(ctx context.Context, manifestNum string, qty float64) error {
	prahariLogger.Info(ctx, "Reported hazardous chemical disposal volume to Environmental Management",
		prahariLogger.String("manifest_num", manifestNum))
	return nil
}
