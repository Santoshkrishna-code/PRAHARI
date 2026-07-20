package export

import (
	"context"
	"fmt"

	"prahari/services/emergency/internal/domain/emergency"
	"prahari/services/emergency/internal/domain/search"
)

type Repository interface {
	SearchEmergencies(ctx context.Context, criteria *search.Criteria) ([]*emergency.Emergency, int64, error)
	GetEmergencyByID(ctx context.Context, id string) (*emergency.Emergency, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	emergencies, _, err := s.repo.SearchEmergencies(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,EmergencyNumber,PlantID,UnitID,Title,Category,Severity,Status,CommanderID,DeclaredAt\n"
	for _, em := range emergencies {
		csvData += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
			em.ID, em.EmergencyNumber, em.PlantID, em.UnitID, em.Title, em.Category, em.Severity, em.Status, em.CommanderID, em.DeclaredAt.Format("2006-01-02 15:04:05"))
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	em, err := s.repo.GetEmergencyByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock Emergency Management Response & Incident Command Report\nID: %s\nNumber: %s\nTitle: %s\nCategory: %s\nSeverity: %s\nStatus: %s\n",
		em.ID, em.EmergencyNumber, em.Title, em.Category, em.Severity, em.Status)
	return []byte(pdfDoc), nil
}
