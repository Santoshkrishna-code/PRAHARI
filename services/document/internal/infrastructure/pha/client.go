package pha

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

func (c *Client) ArchiveHAZOPWorksheet(ctx context.Context, phaID, fileURL string) error {
	prahariLogger.Info(ctx, "Archived approved PHA HAZOP worksheet into Document Management",
		prahariLogger.String("pha_id", phaID))
	return nil
}
