package verification

import (
	"context"
	"fmt"
	"time"

	"prahari/services/moc/internal/domain/changerequest"
	"prahari/services/moc/internal/domain/events"
	"prahari/services/moc/internal/domain/status"
	"prahari/services/moc/internal/domain/verification"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetRequestByID(ctx context.Context, id string) (*changerequest.Request, error)
	SaveRequest(ctx context.Context, req *changerequest.Request) error
	SaveVerification(ctx context.Context, v *verification.Record) error
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

func (s *Service) VerifyChange(ctx context.Context, v *verification.Record) error {
	req, err := s.repo.GetRequestByID(ctx, v.ChangeRequestID)
	if err != nil {
		return err
	}
	v.ID = fmt.Sprintf("ver-%d", time.Now().UnixNano())
	v.VerifiedAt = time.Now()

	if v.PSSRCompleted && v.TrainingVerified && v.DocsUpdated {
		v.Status = "VERIFIED_PASSED"
		if err := status.ValidateTransition(status.Code(req.Status), status.CodeVerification); err == nil {
			req.Status = string(status.CodeVerification)
			req.UpdatedAt = time.Now()
			_ = s.repo.SaveRequest(ctx, req)
			_ = s.publisher.Publish(ctx, events.EventMOCVerified, v)
		}
	} else {
		v.Status = "VERIFIED_FAILED"
	}

	prahariLogger.Info(ctx, "MOC verification record processed", prahariLogger.String("status", v.Status))
	return s.repo.SaveVerification(ctx, v)
}

func (s *Service) CloseoutChange(ctx context.Context, changeRequestID string) error {
	req, err := s.repo.GetRequestByID(ctx, changeRequestID)
	if err != nil {
		return err
	}
	if err := status.ValidateTransition(status.Code(req.Status), status.CodeCloseout); err != nil {
		return err
	}
	req.Status = string(status.CodeCloseout)
	req.UpdatedAt = time.Now()

	if err := s.repo.SaveRequest(ctx, req); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventMOCCclosed, req)
	prahariLogger.Info(ctx, "MOC request closed out successfully", prahariLogger.String("moc_number", req.MOCNumber))
	return nil
}
