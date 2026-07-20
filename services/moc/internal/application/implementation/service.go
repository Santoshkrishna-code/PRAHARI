package implementation

import (
	"context"
	"fmt"
	"time"

	"prahari/services/moc/internal/domain/changerequest"
	"prahari/services/moc/internal/domain/events"
	"prahari/services/moc/internal/domain/implementation"
	"prahari/services/moc/internal/domain/rollback"
	"prahari/services/moc/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetRequestByID(ctx context.Context, id string) (*changerequest.Request, error)
	SaveRequest(ctx context.Context, req *changerequest.Request) error
	SaveImplementation(ctx context.Context, plan *implementation.Plan) error
	SaveRollback(ctx context.Context, plan *rollback.Plan) error
}

type EventPublisher interface {
	Publish(ctx context.Context, eventType string, payload any) error
}

type MaintenanceClient interface {
	CreateWorkOrder(ctx context.Context, plantID, title, description string) (string, error)
}

type Service struct {
	repo        Repository
	publisher   EventPublisher
	maintClient MaintenanceClient
}

func NewService(repo Repository, pub EventPublisher, maintClient MaintenanceClient) *Service {
	return &Service{
		repo:        repo,
		publisher:   pub,
		maintClient: maintClient,
	}
}

func (s *Service) StartImplementation(ctx context.Context, plan *implementation.Plan) error {
	req, err := s.repo.GetRequestByID(ctx, plan.ChangeRequestID)
	if err != nil {
		return err
	}
	if err := status.ValidateTransition(status.Code(req.Status), status.CodeImplementation); err != nil {
		return err
	}
	req.Status = string(status.CodeImplementation)
	req.UpdatedAt = time.Now()

	plan.ID = fmt.Sprintf("imp-%d", time.Now().UnixNano())
	plan.StartDate = time.Now()
	plan.Status = "IN_PROGRESS"

	if s.maintClient != nil && plan.WorkOrderID == "" {
		woID, err := s.maintClient.CreateWorkOrder(ctx, req.PlantID, fmt.Sprintf("MOC Implementation: %s", req.Title), req.Description)
		if err == nil {
			plan.WorkOrderID = woID
		}
	}

	if err := s.repo.SaveImplementation(ctx, plan); err != nil {
		return err
	}

	_ = s.repo.SaveRequest(ctx, req)
	_ = s.publisher.Publish(ctx, events.EventMOCImplemented, plan)
	prahariLogger.Info(ctx, "MOC implementation started", prahariLogger.String("change_request_id", plan.ChangeRequestID))
	return nil
}

func (s *Service) ExecuteRollback(ctx context.Context, rb *rollback.Plan) error {
	req, err := s.repo.GetRequestByID(ctx, rb.ChangeRequestID)
	if err != nil {
		return err
	}
	req.Status = string(status.CodeRolledBack)
	req.UpdatedAt = time.Now()

	rb.ID = fmt.Sprintf("rb-%d", time.Now().UnixNano())
	now := time.Now()
	rb.ExecutedAt = &now
	rb.Status = "EXECUTED"

	if err := s.repo.SaveRollback(ctx, rb); err != nil {
		return err
	}

	_ = s.repo.SaveRequest(ctx, req)
	_ = s.publisher.Publish(ctx, events.EventMOCRollbackExecuted, rb)
	prahariLogger.Warn(ctx, "MOC rollback executed", prahariLogger.String("change_request_id", rb.ChangeRequestID))
	return nil
}
