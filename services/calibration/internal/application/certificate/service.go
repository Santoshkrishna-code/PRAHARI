package certificate

import (
	"context"
	"fmt"
	"time"

	"prahari/services/calibration/internal/domain/calibration"
	"prahari/services/calibration/internal/domain/calibrationcertificate"
	"prahari/services/calibration/internal/domain/events"
	"prahari/services/calibration/internal/domain/status"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetCalibrationByID(ctx context.Context, id string) (*calibration.Record, error)
	SaveCalibration(ctx context.Context, rec *calibration.Record) error
	SaveCertificate(ctx context.Context, cert *calibrationcertificate.Certificate) error
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

func (s *Service) GenerateCertificate(ctx context.Context, calID, certNo string) error {
	rec, err := s.repo.GetCalibrationByID(ctx, calID)
	if err != nil {
		return err
	}

	if err := status.ValidateTransition(status.Code(rec.Status), status.CodeCertificateGenerated); err != nil {
		return err
	}

	cert := &calibrationcertificate.Certificate{
		ID:             fmt.Sprintf("crt-%d", time.Now().UnixNano()),
		CalibrationID:  calID,
		CertificateNo:  certNo,
		IssuedDate:     time.Now(),
		ExpiryDate:     time.Now().Add(365 * 24 * time.Hour),
		DocumentDocRef: "doc-ref-10023",
	}

	rec.Status = string(status.CodeCertificateGenerated)
	rec.CertificateID = cert.ID
	rec.UpdatedAt = time.Now()

	if err := s.repo.SaveCertificate(ctx, cert); err != nil {
		return err
	}
	if err := s.repo.SaveCalibration(ctx, rec); err != nil {
		return err
	}

	_ = s.publisher.Publish(ctx, events.EventCertificateGenerated, cert)
	prahariLogger.Info(ctx, "Calibration certificate generated",
		prahariLogger.String("certificate_no", certNo),
		prahariLogger.String("calibration_id", calID))
	return nil
}
