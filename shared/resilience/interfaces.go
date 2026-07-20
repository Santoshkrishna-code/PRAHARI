package resilience

import (
	"context"
)

// Checker standardizes health checker signatures for third-party systems (databases, APIs, brokers).
type Checker interface {
	Name() string
	Ping(ctx context.Context) error
}
