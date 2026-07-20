package channel

import (
	"context"
)

// Channel defines pluggable adapter interfaces for dispatches.
type Channel interface {
	Send(ctx context.Context, recipient, body string) error
	Name() string
}
