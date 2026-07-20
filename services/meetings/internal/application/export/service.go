package export

import (
	"context"
	"fmt"

	"prahari/services/meetings/internal/domain/meeting"
	"prahari/services/meetings/internal/domain/search"
)

type Repository interface {
	SearchMeetings(ctx context.Context, criteria *search.Criteria) ([]*meeting.Meeting, int64, error)
	GetMeetingByID(ctx context.Context, id string) (*meeting.Meeting, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExportCSV(ctx context.Context, criteria *search.Criteria) ([]byte, error) {
	meetings, _, err := s.repo.SearchMeetings(ctx, criteria)
	if err != nil {
		return nil, err
	}

	csvData := "ID,PlantID,MeetingType,Title,Status,ScheduledAt,OrganizerID\n"
	for _, m := range meetings {
		csvData += fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s\n",
			m.ID, m.PlantID, m.MeetingType, m.Title, m.Status,
			m.ScheduledAt.Format("2006-01-02 15:04:05"), m.OrganizerID)
	}
	return []byte(csvData), nil
}

func (s *Service) ExportPDF(ctx context.Context, id string) ([]byte, error) {
	m, err := s.repo.GetMeetingByID(ctx, id)
	if err != nil {
		return nil, err
	}

	pdfDoc := fmt.Sprintf("%%PDF-1.4 Mock Meeting Minutes Report\nID: %s\nType: %s\nStatus: %s\nTitle: %s\n",
		m.ID, m.MeetingType, m.Status, m.Title)
	return []byte(pdfDoc), nil
}
