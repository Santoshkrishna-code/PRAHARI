package tracking

import (
	"context"
	"fmt"
	"time"

	"prahari/services/vision/internal/domain/tracking"
)

type Repository interface {
	SaveTrackSegment(ctx context.Context, s *tracking.Segment) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) TrackObjectFrame(ctx context.Context, objectID, cameraID string, x, y float64) error {
	seg := &tracking.Segment{
		ID:        fmt.Sprintf("seg-%d", time.Now().UnixNano()),
		ObjectID:  objectID,
		CameraID:  cameraID,
		XVal:      x,
		YVal:      y,
		Timestamp: time.Now(),
	}
	return s.repo.SaveTrackSegment(ctx, seg)
}
