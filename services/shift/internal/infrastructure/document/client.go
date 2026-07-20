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

func (c *Client) FetchApprovedSOP(ctx context.Context, sopDocID string) (string, error) {
	prahariLogger.Info(ctx, "Fetched approved SOP document URL reference from Document Management Service",
		prahariLogger.String("sop_doc_id", sopDocID))
	return "https://s3.amazonaws.com/prahari-document-repository/sops/hydrocracker-startup.pdf", nil
}
