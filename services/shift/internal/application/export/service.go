package export

import (
	"context"
	"fmt"

	"prahari/services/shift/internal/domain/search"
	"prahari/services/shift/internal/domain/shift"
)

type Repository interface {
	SearchShifts(ctx context.Context, criteria *search.Criteria) ([]*shift.Shift, int64, error)
	GetShiftByID(ctx context.Context, id string) (*shift.Shift, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	shifts, _, err := s.repo.SearchShifts(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,ShiftName,PlantID,UnitID,SupervisorID,ScheduledStart,ScheduledEnd,Status\n"
	for _, sh := range shifts {
		csvData += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s\n",
			sh.ID, sh.ShiftName, sh.PlantID, sh.UnitID, sh.SupervisorID, sh.ScheduledStart.Format("2006-01-02 15:04:05"), sh.ScheduledEnd.Format("2006-01-02 15:04:05"), sh.Status)
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	sh, err := s.repo.GetShiftByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock Shift Operations Journal and Handover Protocol Report\nID: %s\nShift: %s\nPlant: %s\nSupervisor: %s\nStatus: %s\n",
		sh.ID, sh.ShiftName, sh.PlantID, sh.SupervisorID, sh.Status)
	return []byte(pdfDoc), nil
}
