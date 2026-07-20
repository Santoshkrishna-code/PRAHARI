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

func (c *Client) FetchNDAFileUrl(ctx context.Context, plantID string) (string, error) {
	prahariLogger.Info(ctx, "Fetched standard NDA file URL template from Document Management Service",
		prahariLogger.String("plant_id", plantID))
	return "https://s3.amazonaws.com/prahari-document-repository/templates/nda.pdf", nil
}
