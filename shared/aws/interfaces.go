package aws

import (
	"context"
)

// HealthChecker is implemented by every AWS service client to report connectivity states.
type HealthChecker interface {
	Ping(ctx context.Context) error
}
