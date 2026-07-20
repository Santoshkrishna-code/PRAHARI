package visualization

import (
	"context"
	"time"

	"prahari/services/digitaltwin/internal/domain/overlay"
	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	SaveOverlay(ctx context.Context, o *overlay.Layer) error
	GetOverlaysByTwin(ctx context.Context, twinID string) ([]*overlay.Layer, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) PushOverlay(ctx context.Context, twinID, layerType, sourceID, label, meta string) (*overlay.Layer, error) {
	lay := &overlay.Layer{
		ID:        "lay-" + time.Now().Format("20060102150405"),
		TwinID:    twinID,
		LayerType: layerType,
		SourceID:  sourceID,
		Label:     label,
		Metadata:  meta,
		Timestamp: time.Now(),
	}

	if err := s.repo.SaveOverlay(ctx, lay); err != nil {
		return nil, err
	}

	prahariLogger.Info(ctx, "Pushed new visualization overlay layer onto Digital Twin canvas",
		prahariLogger.String("twin_id", twinID),
		prahariLogger.String("layer_type", layerType))
	return lay, nil
}

func (s *Service) RenderTwinCanvas(ctx context.Context, twinID string) ([]*overlay.Layer, error) {
	return s.repo.GetOverlaysByTwin(ctx, twinID)
}
