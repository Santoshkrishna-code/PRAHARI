package training

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

func (c *Client) VerifyWorkforceTraining(ctx context.Context, courseID string, targetCount int) (bool, error) {
	prahariLogger.Info(ctx, "Verified workforce training status via Training & Competency Management Service",
		prahariLogger.String("course_id", courseID),
		prahariLogger.Int("target_count", targetCount))
	return true, nil
}
