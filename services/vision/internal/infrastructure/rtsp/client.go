package rtsp

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Client struct {
	url string
}

func NewClient(url string) *Client {
	return &Client{url: url}
}

func (c *Client) PullFrames(ctx context.Context) error {
	prahariLogger.Info(ctx, "Pulling network video stream frames via RTSP protocol",
		prahariLogger.String("url", c.url))
	return nil
}
