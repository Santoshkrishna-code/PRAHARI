package logbook

import (
	"context"
	"fmt"
	"time"

	"prahari/services/shift/internal/domain/operatorjournal"
	"prahari/services/shift/internal/domain/shiftlog"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveLog(ctx context.Context, log *shiftlog.Log) error
	SaveJournal(ctx context.Context, jr *operatorjournal.Journal) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) LogActivity(ctx context.Context, log *shiftlog.Log) error {
	log.ID = fmt.Sprintf("log-%d", time.Now().UnixNano())
	log.Timestamp = time.Now()

	if err := s.repo.SaveLog(ctx, log); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Operator log activity logged",
		prahariLogger.String("shift_id", log.ShiftID),
		prahariLogger.String("category", log.Category))
	return nil
}

func (s *Service) WriteJournal(ctx context.Context, jr *operatorjournal.Journal) error {
	jr.ID = fmt.Sprintf("jnl-%d", time.Now().UnixNano())
	jr.LoggedAt = time.Now()

	if err := s.repo.SaveJournal(ctx, jr); err != nil {
		return err
	}

	prahariLogger.Info(ctx, "Control room operator journal updated", prahariLogger.String("subject", jr.Subject))
	return nil
}
