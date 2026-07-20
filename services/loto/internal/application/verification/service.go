package verification

import (
	"context"
	"fmt"
	"time"

	"prahari/services/loto/internal/domain/events"
	"prahari/services/loto/internal/domain/isolationcertificate"
	"prahari/services/loto/internal/domain/status"
	"prahari/services/loto/internal/domain/verification"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetCertificateByID(ctx context.Context, id string) (*isolationcertificate.Certificate, error)
	SaveCertificate(ctx context.Context, cert *isolationcertificate.Certificate) error
	SaveVerification(ctx context.Context, v *verification.ZeroEnergy) error
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

func (s *Service) VerifyZeroEnergy(ctx context.Context, certID string, v *verification.ZeroEnergy) error {
	cert, err := s.repo.GetCertificateByID(ctx, certID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(cert.Status), status.CodeZeroEnergyVerified); err != nil {
		return err
	}

	v.ID = fmt.Sprintf("ver-%d", time.Now().UnixNano())
	v.CertificateID = certID
	v.VerificationAt = time.Now()

	cert.Status = string(status.CodeZeroEnergyVerified)
	now := time.Now()
	cert.VerifiedAt = &now
	cert.UpdatedAt = time.Now()

	if err := s.repo.SaveVerification(ctx, v); err != nil {
		return err
	}
	if err := s.repo.SaveCertificate(ctx, cert); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventZeroEnergyVerified, cert)
	prahariLogger.Info(ctx, "Zero-energy state verified", prahariLogger.String("certificate_id", certID))
	return nil
}
