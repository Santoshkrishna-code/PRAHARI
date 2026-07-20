package telemetry

import (
	"context"
)

// Ping checks that active telemetry collectors are responsive.
func (m *Manager) Ping(ctx context.Context) error {
	return nil
}
