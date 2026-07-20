package llm

import (
	"context"
	"fmt"

	prahariLogger "prahari/shared/logger"
)

type Client struct {
	model string
}

func NewClient(model string) *Client {
	return &Client{model: model}
}

func (c *Client) Generate(ctx context.Context, prompt string, context []string) (string, error) {
	prahariLogger.Info(ctx, "Executing LLM prompt completion inference request",
		prahariLogger.String("model", c.model))

	// Mock response
	return fmt.Sprintf("Based on the safety logs, this is an automated RCA recommendation output response matching prompt query: %s", prompt), nil
}
