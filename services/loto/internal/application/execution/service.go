package execution

import (
	"context"
	"fmt"
	"time"

	"prahari/services/loto/internal/domain/events"
	"prahari/services/loto/internal/domain/isolationcertificate"
	"prahari/services/loto/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetCertificateByID(ctx context.Context, id string) (*isolationcertificate.Certificate, error)
	SaveCertificate(ctx context.Context, cert *isolationcertificate.Certificate) error
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

func (s *Service) ApproveIsolation(ctx context.Context, cert *isolationcertificate.Certificate) error {
	cert.ID = fmt.Sprintf("lto-%d", time.Now().UnixNano())
	cert.Status = string(status.CodeIsolationApproved)
	cert.CreatedAt = time.Now()
	cert.UpdatedAt = time.Now()

	if err := s.repo.SaveCertificate(ctx, cert); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventIsolationApproved, cert)
	prahariLogger.Info(ctx, "LOTO isolation approved", prahariLogger.String("plan_id", cert.PlanID))
	return nil
}

func (s *Service) ApplyLocksAndTags(ctx context.Context, certID string) error {
	cert, err := s.repo.GetCertificateByID(ctx, certID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(cert.Status), status.CodeLocksApplied); err != nil {
		return err
	}
	cert.Status = string(status.CodeLocksApplied)
	cert.UpdatedAt = time.Now()

	if err := s.repo.SaveCertificate(ctx, cert); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventLocksApplied, cert)
	prahariLogger.Info(ctx, "LOTO physical locks applied to points", prahariLogger.String("certificate_id", certID))
	return nil
}
