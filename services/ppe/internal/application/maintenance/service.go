package maintenance

import (
	"context"
	"fmt"
	"time"

	"prahari/services/ppe/internal/domain/events"
	"prahari/services/ppe/internal/domain/ppeitem"
	"prahari/services/ppe/internal/domain/ppemaintenance"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetPPEItemByID(ctx context.Context, id string) (*ppeitem.Item, error)
	SavePPEItem(ctx context.Context, item *ppeitem.Item) error
	SavePPEMaintenance(ctx context.Context, rec *ppemaintenance.Record) error
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

func (s *Service) MaintainPPEItem(ctx context.Context, rec *ppemaintenance.Record) error {
	item, err := s.repo.GetPPEItemByID(ctx, rec.ItemID)
	if err != nil {
		return err
	}

	rec.ID = fmt.Sprintf("mnt-%d", time.Now().UnixNano())
	rec.MaintenanceAt = time.Now()

	item.Status = "MAINTENANCE"

	if err := s.repo.SavePPEMaintenance(ctx, rec); err != nil {
		return err
	}
	if err := s.repo.SavePPEItem(ctx, item); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventPPEMaintenanceCompleted, rec)
	prahariLogger.Info(ctx, "PPE protective gear sent for technical servicing/cleaning",
		prahariLogger.String("item_id", rec.ItemID))
	return nil
}
