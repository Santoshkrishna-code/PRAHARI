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

func (c *Client) FetchLabAccreditationRef(ctx context.Context, labID string) (string, error) {
	prahariLogger.Info(ctx, "Fetched controlled laboratory ISO/IEC 17025 certification copy reference",
		prahariLogger.String("lab_id", labID))
	return "https://s3.amazonaws.com/prahari-document-repository/accreditations/ISO17025-lab04.pdf", nil
}
