package sqs

import (
	"context"
	"fmt"
)

// RedirectToDLQ routes a poisoned, unprocessable message body into the designated dead-letter queue.
func (c *Client) RedirectToDLQ(ctx context.Context, dlqURL, messageBody string) (string, error) {
	msgID, err := c.SendMessage(ctx, dlqURL, messageBody)
	if err != nil {
		return "", fmt.Errorf("failed to redirect message to DLQ: %w", err)
	}
	return msgID, nil
}
