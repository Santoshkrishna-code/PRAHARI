package xray

import (
	"context"
	"fmt"
)

// Segment wraps an active trace segment span.
type Segment struct {
	name string
}

// Client wraps X-Ray operations.
type Client struct {
	enabled bool
}

// NewClient constructs an X-Ray client wrapper.
func NewClient(enabled bool) *Client {
	return &Client{enabled: enabled}
}

// Ping implements aws.HealthChecker checking connectivity.
func (c *Client) Ping(ctx context.Context) error {
	return nil // X-Ray is a daemon-based receiver: returns nil to indicate active state
}

// StartSegment starts a new trace subsegment.
func (c *Client) StartSegment(ctx context.Context, name string) (context.Context, *Segment) {
	if !c.enabled {
		return ctx, &Segment{name: name}
	}

	// In production, integrate:
	// ctx, segment := xray.BeginSubsegment(ctx, name)
	// return ctx, &Segment{name: name}

	fmt.Printf("[XRAY-MOCK] Started segment: %s\n", name)
	return ctx, &Segment{name: name}
}

// Close closes the tracing segment.
func (s *Segment) Close() {
	fmt.Printf("[XRAY-MOCK] Closed segment: %s\n", s.name)
}
