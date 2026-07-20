package locks

import (
	"context"
	"time"
)

// StartLeaseExtension runs a background watchdog that regularly refreshes the lock TTL.
// Returns a cancel function to stop the extension loop when the task completes.
func StartLeaseExtension(ctx context.Context, m *Mutex, interval, extension time.Duration) context.CancelFunc {
	childCtx, cancel := context.WithCancel(ctx)
	
	// Lua script: extend TTL only if key holds our unique value
	const extendScript = `
		if redis.call("get", KEYS[1]) == ARGV[1] then
			return redis.call("pexpire", KEYS[1], ARGV[2])
		else
			return 0
		end
	`

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-childCtx.Done():
				return
			case <-ticker.C:
				millis := extension.Milliseconds()
				_, _ = m.client.UniversalClient.Eval(
					childCtx,
					extendScript,
					[]string{m.key},
					m.value,
					millis,
				).Result()
			}
		}
	}()

	return cancel
}
