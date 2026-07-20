package export

import (
	"context"
	"fmt"
	"prahari/services/water/internal/domain/search"
	"prahari/services/water/internal/domain/waterprofile"
)

type Repository interface {
	SearchProfiles(ctx context.Context, criteria *search.Criteria) ([]*waterprofile.Profile, int64, error)
	GetProfileByID(ctx context.Context, id string) (*waterprofile.Profile, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	profiles, _, err := s.repo.SearchProfiles(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,PlantID,FacilityName,WaterBasinRegion,AnnualBudgetKL,TargetRecyclePct,Status\n"
	for _, p := range profiles {
		csvData += fmt.Sprintf("%s,%s,%s,%s,%.2f,%.2f,%s\n",
			p.ID, p.PlantID, p.FacilityName, p.WaterBasinRegion, p.AnnualBudgetKL, p.TargetRecyclePct, p.Status)
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	profile, err := s.repo.GetProfileByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock Water Profile Executive Report\nID: %s\nFacility: %s\nBasin: %s\nBudget: %.2f KL\n",
		profile.ID, profile.FacilityName, profile.WaterBasinRegion, profile.AnnualBudgetKL)
	return []byte(pdfDoc), nil
}
