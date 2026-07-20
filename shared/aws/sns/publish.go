package sns

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

// Client wraps the SNS SDK client.
type Client struct {
	client *sns.Client
}

// NewClient constructs an SNS client wrapper.
func NewClient(client *sns.Client) *Client {
	return &Client{client: client}
}

// Ping implements aws.HealthChecker checking connectivity by listing topics.
func (c *Client) Ping(ctx context.Context) error {
	if c.client == nil {
		return fmt.Errorf("sns client is uninitialized")
	}
	_, err := c.client.ListTopics(ctx, &sns.ListTopicsInput{})
	return err
}

// PublishMessage sends a raw message payload to an SNS topic ARN.
func (c *Client) PublishMessage(ctx context.Context, topicARN, message string) (string, error) {
	if c.client == nil {
		return "", fmt.Errorf("sns client is uninitialized")
	}

	output, err := c.client.Publish(ctx, &sns.PublishInput{
		TopicArn: &topicARN,
		Message:  &message,
	})
	if err != nil {
		return "", fmt.Errorf("failed to publish to SNS topic %s: %w", topicARN, err)
	}

	return *output.MessageId, nil
}

// PublishSMS dispatches direct SMS notifications to target mobile numbers.
func (c *Client) PublishSMS(ctx context.Context, phoneNumber, message string) (string, error) {
	if c.client == nil {
		return "", fmt.Errorf("sns client is uninitialized")
	}

	output, err := c.client.Publish(ctx, &sns.PublishInput{
		PhoneNumber: &phoneNumber,
		Message:     &message,
	})
	if err != nil {
		return "", fmt.Errorf("failed to send SMS notification: %w", err)
	}

	return *output.MessageId, nil
}
