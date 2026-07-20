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

func (c *Client) TriggerPHASignoffWorkflow(ctx context.Context, studyID string) error {
	prahariLogger.Info(ctx, "Triggered PHA study sign-off workflow in Workflow Engine", prahariLogger.String("study_id", studyID))
	return nil
}
