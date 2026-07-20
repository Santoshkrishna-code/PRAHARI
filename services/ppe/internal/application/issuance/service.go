package issuance

import (
	"context"
	"fmt"
	"time"

	"prahari/services/ppe/internal/domain/events"
	"prahari/services/ppe/internal/domain/ppeitem"
	"prahari/services/ppe/internal/domain/ppeissue"
	prahariLogger "prahari/shared/logger"
)


type Repository interface {
	GetPPEItemByID(ctx context.Context, id string) (*ppeitem.Item, error)
	SavePPEItem(ctx context.Context, item *ppeitem.Item) error
	SavePPEIssue(ctx context.Context, issue *ppeissue.Record) error
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

func (s *Service) IssuePPEItem(ctx context.Context, rec *ppeissue.Record) error {
	item, err := s.repo.GetPPEItemByID(ctx, rec.ItemID)
	if err != nil {
		return err
	}

	rec.ID = fmt.Sprintf("iss-%d", time.Now().UnixNano())
	rec.IssuedAt = time.Now()
	rec.ExpectedReturn = time.Now().Add(8 * time.Hour)

	item.Status = "ISSUED"
	item.IssuedTo = rec.IssuedToID

	if err := s.repo.SavePPEIssue(ctx, rec); err != nil {
		return err
	}
	if err := s.repo.SavePPEItem(ctx, item); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventPPEIssued, rec)
	prahariLogger.Info(ctx, "PPE protective gear issued to user",
		prahariLogger.String("item_id", rec.ItemID),
		prahariLogger.String("user_id", rec.IssuedToID))
	return nil
}
