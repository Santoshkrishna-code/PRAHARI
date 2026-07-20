package polling

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// PollEventsLoop reads partition offsets in the background, executing callbacks.
func PollEventsLoop(ctx context.Context) error {
	prahariLogger.Info(ctx, "Starting Kafka background events poll loop...")
	return nil
}
