package vectorstore

import (
	"context"

	"prahari/services/ai/internal/domain/retrieval"
	prahariLogger "prahari/shared/logger"
)

type Client struct {
	addr string
}

func NewClient(addr string) *Client {
	return &Client{addr: addr}
}

func (c *Client) IndexChunk(ctx context.Context, chunkID string, embedding []float32) error {
	prahariLogger.Info(ctx, "Indexed text chunk float32 array vector into vector database",
		prahariLogger.String("chunk_id", chunkID))
	return nil
}

func (c *Client) QuerySimilar(ctx context.Context, embedding []float32, limit int) ([]*retrieval.Result, error) {
	prahariLogger.Info(ctx, "Queried vector similarity nearest neighbors matches")
	return []*retrieval.Result{
		{ChunkID: "chk-mock-01", DocID: "doc-mock-01", Text: "Banned chemicals and SDS safety criteria procedures", Confidence: 0.94},
	}, nil
}
