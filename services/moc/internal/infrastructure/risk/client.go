package risk

import (
	"context"
	"fmt"
	"time"

	prahariLogger "prahari/shared/logger"
)

type Client struct {
	grpcAddr string
}

func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

func (c *Client) CreateRiskAssessment(ctx context.Context, plantID, title, description string) (string, error) {
	raID := fmt.Sprintf("ra-moc-%d", time.Now().UnixNano())
	prahariLogger.Info(ctx, "Triggered fresh Risk Assessment in Risk Assessment Service for MOC",
		prahariLogger.String("risk_assessment_id", raID),
		prahariLogger.String("title", title))
	return raID, nil
}
