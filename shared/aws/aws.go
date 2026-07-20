package aws

import (
	"context"
)

// Manager aggregates individual mockable client wrappers for AWS Cloud services.
type Manager struct {
	cfg Config
}

// NewManager constructs a consolidated AWS manager.
func NewManager(cfg Config) *Manager {
	return &Manager{
		cfg: cfg,
	}
}

// Health checks states across registered services.
func (m *Manager) Health(ctx context.Context, checkers map[string]HealthChecker) map[string]error {
	return CheckAWSHealth(ctx, checkers)
}
