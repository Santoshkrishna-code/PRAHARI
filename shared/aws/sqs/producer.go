package sqs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// Client wraps the AWS SQS SDK client.
type Client struct {
	client *sqs.Client
}

// NewClient constructs an SQS client wrapper.
func NewClient(client *sqs.Client) *Client {
	return &Client{client: client}
}

// Ping implements aws.HealthChecker checking connectivity by listing queues.
func (c *Client) Ping(ctx context.Context) error {
	if c.client == nil {
		return fmt.Errorf("sqs client is uninitialized")
	}
	_, err := c.client.ListQueues(ctx, &sqs.ListQueuesInput{})
	return err
}

// SendMessage puts a string payload into the target SQS queue.
func (c *Client) SendMessage(ctx context.Context, queueURL, body string) (string, error) {
	return c.SendMessageWithDelay(ctx, queueURL, body, 0)
}

// SendMessageWithDelay puts a message payload with delay offsets.
func (c *Client) SendMessageWithDelay(ctx context.Context, queueURL, body string, delaySeconds int32) (string, error) {
	if c.client == nil {
		return "", fmt.Errorf("sqs client is uninitialized")
	}

	input := &sqs.SendMessageInput{
		QueueUrl:     &queueURL,
		MessageBody:  &body,
		DelaySeconds: delaySeconds,
	}

	output, err := c.client.SendMessage(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to send SQS message to %s: %w", queueURL, err)
	}

	return *output.MessageId, nil
}
