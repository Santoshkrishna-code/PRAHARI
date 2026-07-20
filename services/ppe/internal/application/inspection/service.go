package inspection

import (
	"context"
	"fmt"
	"time"

	"prahari/services/ppe/internal/domain/events"
	"prahari/services/ppe/internal/domain/ppeinspection"
	"prahari/services/ppe/internal/domain/ppeitem"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetPPEItemByID(ctx context.Context, id string) (*ppeitem.Item, error)
	SavePPEItem(ctx context.Context, item *ppeitem.Item) error
	SavePPEInspection(ctx context.Context, rec *ppeinspection.Record) error
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

func (s *Service) InspectPPEItem(ctx context.Context, rec *ppeinspection.Record) error {
	item, err := s.repo.GetPPEItemByID(ctx, rec.ItemID)
	if err != nil {
		return err
	}

	rec.ID = fmt.Sprintf("ins-%d", time.Now().UnixNano())
	rec.InspectedAt = time.Now()

	now := time.Now()
	item.LastInspectedAt = &now
	if rec.Result == "PASS" {
		item.Status = "AVAILABLE"
	} else {
		item.Status = "DISPOSED" // or maintenance
	}

	if err := s.repo.SavePPEInspection(ctx, rec); err != nil {
		return err
	}
	if err := s.repo.SavePPEItem(ctx, item); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventPPEInspected, rec)
	prahariLogger.Info(ctx, "PPE protective gear inspected",
		prahariLogger.String("item_id", rec.ItemID),
		prahariLogger.String("result", rec.Result))
	return nil
}
