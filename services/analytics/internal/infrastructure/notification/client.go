package notification

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

func (c *Client) DistributePDFReport(ctx context.Context, title string, recipients []string) error {
	prahariLogger.Info(ctx, "Scheduled distribution of compiled analytical PDF report",
		prahariLogger.String("title", title))
	return nil
}
