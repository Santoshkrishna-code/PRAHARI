package workflow

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

func (c *Client) SubmitReportApproval(ctx context.Context, reportID string) error {
	prahariLogger.Info(ctx, "Submitting corporate sustainability report to Workflow Engine for executive signoff routing",
		prahariLogger.String("report_id", reportID))
	return nil
}
