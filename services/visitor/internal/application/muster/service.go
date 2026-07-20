package muster

import (
	"context"
	"fmt"
	"time"

	"prahari/services/visitor/internal/domain/events"
	"prahari/services/visitor/internal/domain/emergencymuster"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveMusterRecord(ctx context.Context, rec *emergencymuster.Record) error
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type Service struct {
	repo      Repository
	publisher EventPublisher
}

func NewService(repo Repository, pub EventPublisher) *Service {
	return &Service{
		repo:      repo,
		publisher: pub,
	}
}

func (s *Service) AccountForVisitor(ctx context.Context, rec *emergencymuster.Record) error {
	rec.ID = fmt.Sprintf("mst-%d", time.Now().UnixNano())
	now := time.Now()
	rec.AccountedAt = &now
	rec.AccountedFor = true

	if err := s.repo.SaveMusterRecord(ctx, rec); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventVisitorMusterComplete, rec)
	prahariLogger.Info(ctx, "Evacuated visitor accounted for at assembly point muster check",
		prahariLogger.String("visitor_id", rec.VisitorID))
	return nil
}
