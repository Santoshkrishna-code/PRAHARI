package export

import (
	"context"
	"fmt"

	"prahari/services/calibration/internal/domain/calibration"
	"prahari/services/calibration/internal/domain/search"
)

type Repository interface {
	SearchCalibrations(ctx context.Context, criteria *search.Criteria) ([]*calibration.Record, int64, error)
	GetCalibrationByID(ctx context.Context, id string) (*calibration.Record, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	calibrations, _, err := s.repo.SearchCalibrations(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,InstrumentID,CalibratedBy,CalibratedAt,Status,Result,CertificateID\n"
	for _, c := range calibrations {
		csvData += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s\n",
			c.ID, c.InstrumentID, c.CalibratedBy, c.CalibratedAt.Format("2006-01-02 15:04:05"), c.Status, c.Result, c.CertificateID)
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	c, err := s.repo.GetCalibrationByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock Instrument Metrology Verification & ISO 17025 Traceability Calibration Certificate Report\nID: %s\nInstrument: %s\nStatus: %s\nResult: %s\nApproved By: %s\n",
		c.ID, c.InstrumentID, c.Status, c.Result, c.ApprovedBy)
	return []byte(pdfDoc), nil
}
