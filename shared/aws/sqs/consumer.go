package sqs

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

// StartConsumer runs a blocking polling loop retrieving messages and dispatching them to the handler callback.
// If the handler completes without errors, the message is automatically deleted from SQS.
func (c *Client) StartConsumer(ctx context.Context, queueURL string, handler func(ctx context.Context, body string) error) error {
	if c.client == nil {
		return fmt.Errorf("sqs client is uninitialized")
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// Poll SQS (Long polling enabled)
			input := &sqs.ReceiveMessageInput{
				QueueUrl:            &queueURL,
				MaxNumberOfMessages: 10,
				WaitTimeSeconds:     20, // 20s long polling to optimize network call costs
				VisibilityTimeout:   30,
			}

			output, err := c.client.ReceiveMessage(ctx, input)
			if err != nil {
				// Prevent loop thrashing on network errors: backoff before retrying
				select {
				case <-ctx.Done():
					return ctx.Err()
				case <-time.After(5 * time.Second):
					continue
				}
			}

			for _, msg := range output.Messages {
				if msg.Body == nil {
					continue
				}

				// Execute handler callback with correlation tracing contexts
				err := handler(ctx, *msg.Body)
				if err == nil {
					// Successful execution: delete message
					_, _ = c.client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
						QueueUrl:      &queueURL,
						ReceiptHandle: msg.ReceiptHandle,
					})
				}
			}
		}
	}
}

// DeleteMessage manually deletes an SQS message if callers bypass automatic consumers.
func (c *Client) DeleteMessage(ctx context.Context, queueURL, receiptHandle string) error {
	if c.client == nil {
		return fmt.Errorf("sqs client is uninitialized")
	}
	_, err := c.client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      &queueURL,
		ReceiptHandle: &receiptHandle,
	})
	return err
}
