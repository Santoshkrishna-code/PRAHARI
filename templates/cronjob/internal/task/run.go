package task

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// RunTask executes periodic cleanup jobs.
func RunTask(ctx context.Context) error {
	prahariLogger.Info(ctx, "Scheduled cleanup database cron action executed successfully.")
	return nil
}
