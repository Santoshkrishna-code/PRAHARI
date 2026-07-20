package embeddings

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Client struct {
	model string
}

func NewClient(model string) *Client {
	return &Client{model: model}
}

func (c *Client) GenerateEmbedding(ctx context.Context, text string) ([]float32, error) {
	prahariLogger.Info(ctx, "Generated document chunk vector embeddings float32 dimensions")
	// return mock 3-dim vector
	return []float32{0.1, 0.42, -0.99}, nil
}
