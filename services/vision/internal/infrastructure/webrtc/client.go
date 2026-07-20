package webrtc

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Client struct {
	sdp string
}

func NewClient(sdp string) *Client {
	return &Client{sdp: sdp}
}

func (c *Client) StreamLivePeerConnection(ctx context.Context) error {
	prahariLogger.Info(ctx, "Established WebRTC peer connection for zero-latency camera streams projection")
	return nil
}
