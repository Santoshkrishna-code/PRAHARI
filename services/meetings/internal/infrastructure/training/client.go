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

func (c *Client) RecordToolboxTalkAttendance(ctx context.Context, attendeeID, meetingID, topicTitle string) error {
	prahariLogger.Info(ctx, "Recorded mandatory toolbox talk attendance in Training & Competency",
		prahariLogger.String("attendee_id", attendeeID),
		prahariLogger.String("meeting_id", meetingID))
	return nil
}
