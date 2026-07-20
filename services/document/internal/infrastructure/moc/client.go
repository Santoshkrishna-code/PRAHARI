package moc

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

func (c *Client) TriggerDocumentRevisionForMOC(ctx context.Context, mocID, documentID string) error {
	prahariLogger.Info(ctx, "Triggered document version revision request from approved MOC",
		prahariLogger.String("moc_id", mocID),
		prahariLogger.String("document_id", documentID))
	return nil
}
