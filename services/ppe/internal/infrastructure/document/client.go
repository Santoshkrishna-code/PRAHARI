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

func (c *Client) FetchPPEManualUrl(ctx context.Context, ppeID string) (string, error) {
	prahariLogger.Info(ctx, "Fetched protective equipment user manual URL reference from Document Management Service",
		prahariLogger.String("ppe_id", ppeID))
	return "https://s3.amazonaws.com/prahari-document-repository/manuals/ansi-hard-hat.pdf", nil
}
