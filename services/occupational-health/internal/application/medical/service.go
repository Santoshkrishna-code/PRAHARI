package medical

import (
	"context"
	"time"

	"prahari/services/occupational-health/internal/domain/appointment"
	"prahari/services/occupational-health/internal/domain/clinic"
	"prahari/services/occupational-health/internal/domain/laboratory"
	"prahari/services/occupational-health/internal/domain/laboratoryresult"
	"prahari/services/occupational-health/internal/domain/medicalexamination"
	"prahari/services/occupational-health/internal/domain/medicalrecord"
	"prahari/services/occupational-health/internal/domain/physician"
)

// Repository defines ports for medical records data access.
type Repository interface {
	SaveRecord(ctx context.Context, r *medicalrecord.MedicalRecord) error
	SaveExam(ctx context.Context, e *medicalexamination.MedicalExamination) error
	SaveAppointment(ctx context.Context, a *appointment.Appointment) error
	SavePhysician(ctx context.Context, p *physician.Physician) error
	SaveClinic(ctx context.Context, c *clinic.Clinic) error
	SaveLaboratory(ctx context.Context, l *laboratory.Laboratory) error
	SaveLaboratoryResult(ctx context.Context, lr *laboratoryresult.LaboratoryResult) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateMedicalRecord(ctx context.Context, r *medicalrecord.MedicalRecord) error {
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
	if err := r.Validate(); err != nil {
		return err
	}
	return s.repo.SaveRecord(ctx, r)
}

func (s *Service) RecordExamination(ctx context.Context, e *medicalexamination.MedicalExamination) error {
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	if err := e.Validate(); err != nil {
		return err
	}
	return s.repo.SaveExam(ctx, e)
}

func (s *Service) ScheduleAppointment(ctx context.Context, a *appointment.Appointment) error {
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	if err := a.Validate(); err != nil {
		return err
	}
	return s.repo.SaveAppointment(ctx, a)
}
