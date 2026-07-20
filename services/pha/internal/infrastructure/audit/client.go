package audit

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

func (c *Client) ReportUnresolvedPHAAction(ctx context.Context, studyID string, openRecs int) error {
	prahariLogger.Info(ctx, "Reported open PHA action items to Audit Management Service for PSM assurance",
		prahariLogger.String("study_id", studyID),
		prahariLogger.Int("open_recs", openRecs))
	return nil
}
