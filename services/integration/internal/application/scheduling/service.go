package scheduling

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) StartScheduler(ctx context.Context) {
	prahariLogger.Info(ctx, "Starting Integration Jobs Cron Scheduler")
}
