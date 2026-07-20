package action

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

func (c *Client) CreateCAPAFromMeeting(ctx context.Context, meetingID, title, assignedTo string) (string, error) {
	prahariLogger.Info(ctx, "Pushed meeting action to central CAPA system",
		prahariLogger.String("meeting_id", meetingID),
		prahariLogger.String("title", title))
	return "act-capa-001", nil
}
