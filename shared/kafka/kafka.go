package kafka

import (
	"context"
)

// Manager coordinates configuration handles and diagnostic checks.
type Manager struct {
	cfg Config
}

// NewManager constructs a new Manager instance.
func NewManager(cfg Config) *Manager {
	return &Manager{cfg: cfg}
}

// Ping checks broker clusters connectivity.
func (m *Manager) Ping(ctx context.Context) error {
	return Ping(ctx, m.cfg)
}
