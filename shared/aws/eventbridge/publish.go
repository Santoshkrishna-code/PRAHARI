package eventbridge

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
)

// Client wraps the AWS EventBridge SDK client.
type Client struct {
	client *eventbridge.Client
}

// NewClient constructs an EventBridge client wrapper.
func NewClient(client *eventbridge.Client) *Client {
	return &Client{client: client}
}

// Ping implements aws.HealthChecker checking connectivity by listing event buses.
func (c *Client) Ping(ctx context.Context) error {
	if c.client == nil {
		return fmt.Errorf("eventbridge client is uninitialized")
	}
	_, err := c.client.ListEventBuses(ctx, &eventbridge.ListEventBusesInput{Limit: &[]int32{1}[0]})
	return err
}

// PublishEvent puts a single custom event payload onto the specified event bus.
func (c *Client) PublishEvent(ctx context.Context, busName, source, detailType, detailJSON string) (string, error) {
	if c.client == nil {
		return "", fmt.Errorf("eventbridge client is uninitialized")
	}

	entry := types.PutEventsRequestEntry{
		EventBusName: &busName,
		Source:       &source,
		DetailType:   &detailType,
		Detail:       &detailJSON,
	}

	input := &eventbridge.PutEventsInput{
		Entries: []types.PutEventsRequestEntry{entry},
	}

	output, err := c.client.PutEvents(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to put events to bus %s: %w", busName, err)
	}

	if output.FailedEntryCount > 0 && len(output.Entries) > 0 {
		return "", fmt.Errorf("event publication failed with error: %s", *output.Entries[0].ErrorMessage)
	}

	if len(output.Entries) == 0 || output.Entries[0].EventId == nil {
		return "", fmt.Errorf("received empty response entry from EventBridge")
	}

	return *output.Entries[0].EventId, nil
}
