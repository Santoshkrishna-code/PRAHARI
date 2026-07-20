package inspection

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

func (c *Client) FetchInspectionFindings(ctx context.Context, inspectionID string) ([]string, error) {
	prahariLogger.Info(ctx, "Fetched inspection findings to raise corrective actions",
		prahariLogger.String("inspection_id", inspectionID))
	return []string{"FND-01: corroded pipeline clamp", "FND-02: blocked fire extinguisher"}, nil
}
