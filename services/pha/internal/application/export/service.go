package export

import (
	"context"
	"fmt"

	"prahari/services/pha/internal/domain/phastudy"
	"prahari/services/pha/internal/domain/search"
)

type Repository interface {
	SearchStudies(ctx context.Context, criteria *search.Criteria) ([]*phastudy.Study, int64, error)
	GetStudyByID(ctx context.Context, id string) (*phastudy.Study, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	studies, _, err := s.repo.SearchStudies(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,StudyNumber,PlantID,UnitID,Title,Method,Status,LeaderID\n"
	for _, st := range studies {
		csvData += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s\n",
			st.ID, st.StudyNumber, st.PlantID, st.UnitID, st.Title, st.Method, st.Status, st.LeaderID)
	}

	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	st, err := s.repo.GetStudyByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock Process Hazard Analysis Executive Report\nID: %s\nStudy Number: %s\nTitle: %s\nMethod: %s\nStatus: %s\n",
		st.ID, st.StudyNumber, st.Title, st.Method, st.Status)
	return []byte(pdfDoc), nil
}
