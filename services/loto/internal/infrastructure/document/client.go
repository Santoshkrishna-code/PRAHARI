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

func (c *Client) FetchDeisolationProcedureRef(ctx context.Context, planID string) (string, error) {
	prahariLogger.Info(ctx, "Fetched controlled de-isolation restoration standard operating procedure reference link",
		prahariLogger.String("plan_id", planID))
	return "https://s3.amazonaws.com/prahari-document-repository/procedures/deisolation-restoration-proc.pdf", nil
}
