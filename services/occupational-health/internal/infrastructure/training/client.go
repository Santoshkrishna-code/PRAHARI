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

func (c *Client) TriggerRefresherTraining(ctx context.Context, workerID string, courseName string) error {
	prahariLogger.Info(ctx, "Triggering mandatory return-to-work safety refresher induction via Training Service gRPC",
		prahariLogger.String("worker_id", workerID),
		prahariLogger.String("course_name", courseName))
	return nil
}
