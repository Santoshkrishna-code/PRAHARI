package workorder

import (
	"context"

	workorderDomain "prahari/services/maintenance/internal/domain/workorder"
)

// Repository query work order persistence stores.
type Repository interface {
	Create(ctx context.Context, w *workorderDomain.WorkOrder) error
	FindByID(ctx context.Context, id string) (*workorderDomain.WorkOrder, error)
	FindByMaintenanceID(ctx context.Context, maintenanceID string) ([]*workorderDomain.WorkOrder, error)
}

// Service manages workorder executions.
type Service struct {
	repo Repository
}

// NewService instantiates Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// GetWorkOrders returns assigned walkthrough cards.
func (s *Service) GetWorkOrders(ctx context.Context, maintenanceID string) ([]*workorderDomain.WorkOrder, error) {
	return s.repo.FindByMaintenanceID(ctx, maintenanceID)
}
