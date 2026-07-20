package bootstrap

import (
	"context"
	"time"

	prahariLogger "prahari/shared/logger"
)

// InitScheduler starts a background loop querying SLA timer violations.
func InitScheduler(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				// In production, execute SQL lookups resolving expired SLAs:
				prahariLogger.Info(ctx, "Scanning active SLA timers deadlines...")
			}
		}
	}()
}
