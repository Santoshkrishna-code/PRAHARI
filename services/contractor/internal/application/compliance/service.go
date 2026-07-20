package compliance

import (
	"context"

	certDomain "prahari/services/contractor/internal/domain/certification"
	trainingDomain "prahari/services/contractor/internal/domain/training"
)

// CertRepository query certifications records.
type CertRepository interface {
	Create(ctx context.Context, c *certDomain.Certification) error
	FindByWorkerID(ctx context.Context, workerID string) ([]*certDomain.Certification, error)
}

// TrainingRepository query training records.
type TrainingRepository interface {
	Create(ctx context.Context, t *trainingDomain.Training) error
	FindByWorkerID(ctx context.Context, workerID string) ([]*trainingDomain.Training, error)
}

// Service manages compliance credentials.
type Service struct {
	certRepo     CertRepository
	trainingRepo TrainingRepository
}

// NewService instantiates Compliance Service.
func NewService(certRepo CertRepository, trainingRepo TrainingRepository) *Service {
	return &Service{
		certRepo:     certRepo,
		trainingRepo: trainingRepo,
	}
}

// VerifyCompliance checks certifications and training expiry.
func (s *Service) VerifyCompliance(ctx context.Context, workerID string) (bool, error) {
	certs, _ := s.certRepo.FindByWorkerID(ctx, workerID)
	for _, cert := range certs {
		if cert.IsExpired() {
			return false, nil
		}
	}
	return true, nil
}
