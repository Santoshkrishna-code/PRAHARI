package jwt

import (
	"context"
	"time"
)

// StartRotationLoop launches a background goroutine to periodically refresh the JWKS key cache.
func StartRotationLoop(ctx context.Context, resolver *KeyResolver, interval time.Duration, logErr func(err error)) {
	ticker := time.NewTicker(interval)
	
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				// Refresh the keys. Ignores errors to prevent crashes (retains stale keys on network fails)
				if err := resolver.refreshKeys(ctx); err != nil && logErr != nil {
					logErr(err)
				}
			}
		}
	}()
}
