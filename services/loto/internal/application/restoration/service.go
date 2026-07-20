package restoration

import (
	"context"
	"fmt"
	"time"

	"prahari/services/loto/internal/domain/events"
	"prahari/services/loto/internal/domain/isolationcertificate"
	"prahari/services/loto/internal/domain/restoration"
	"prahari/services/loto/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetCertificateByID(ctx context.Context, id string) (*isolationcertificate.Certificate, error)
	SaveCertificate(ctx context.Context, cert *isolationcertificate.Certificate) error
	SaveRestoration(ctx context.Context, r *restoration.Record) error
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

func (s *Service) RestoreSystem(ctx context.Context, certID string, r *restoration.Record) error {
	cert, err := s.repo.GetCertificateByID(ctx, certID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(cert.Status), status.CodeReturnedToService); err != nil {
		return err
	}

	r.ID = fmt.Sprintf("rst-%d", time.Now().UnixNano())
	r.CertificateID = certID
	r.RestoredAt = time.Now()

	cert.Status = string(status.CodeReturnedToService)
	now := time.Now()
	cert.RestoredAt = &now
	cert.UpdatedAt = time.Now()

	if err := s.repo.SaveRestoration(ctx, r); err != nil {
		return err
	}
	if err := s.repo.SaveCertificate(ctx, cert); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventSystemRestored, cert)
	prahariLogger.Info(ctx, "System restored to operation", prahariLogger.String("certificate_id", certID))
	return nil
}
