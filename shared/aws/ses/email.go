package ses

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

// Client wraps the AWS SES SDK client.
type Client struct {
	client *ses.Client
}

// NewClient constructs an SES client wrapper.
func NewClient(client *ses.Client) *Client {
	return &Client{client: client}
}

// Ping implements aws.HealthChecker checking connectivity by retrieving send quota.
func (c *Client) Ping(ctx context.Context) error {
	if c.client == nil {
		return fmt.Errorf("ses client is uninitialized")
	}
	_, err := c.client.GetSendQuota(ctx, &ses.GetSendQuotaInput{})
	return err
}

// SendEmail dispatches an HTML email payload to a single destination address.
func (c *Client) SendEmail(ctx context.Context, from, to, subject, htmlBody string) (string, error) {
	if c.client == nil {
		return "", fmt.Errorf("ses client is uninitialized")
	}

	input := &ses.SendEmailInput{
		Source: &from,
		Destination: &types.Destination{
			ToAddresses: []string{to},
		},
		Message: &types.Message{
			Subject: &types.Content{
				Data: &subject,
			},
			Body: &types.Body{
				Html: &types.Content{
					Data: &htmlBody,
				},
			},
		},
	}

	output, err := c.client.SendEmail(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to send SES email from %s to %s: %w", from, to, err)
	}

	return *output.MessageId, nil
}
