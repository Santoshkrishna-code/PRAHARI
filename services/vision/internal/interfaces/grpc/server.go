package grpc

import (
	"context"

	inferenceApp "prahari/services/vision/internal/application/inference"
	searchApp "prahari/services/vision/internal/application/search"
	streamingApp "prahari/services/vision/internal/application/streaming"
	"prahari/services/vision/internal/domain/camera"
	"prahari/services/vision/internal/domain/detection"
	"prahari/services/vision/internal/domain/inference"
	"prahari/services/vision/internal/domain/search"
)

type Server struct {
	streamingSvc *streamingApp.Service
	inferenceSvc *inferenceApp.Service
	searchSvc    *searchApp.Service
}

func NewServer(
	streamingSvc *streamingApp.Service,
	inferenceSvc *inferenceApp.Service,
	searchSvc *searchApp.Service,
) *Server {
	return &Server{
		streamingSvc: streamingSvc,
		inferenceSvc: inferenceSvc,
		searchSvc:    searchSvc,
	}
}

func (s *Server) RegisterCamera(ctx context.Context, c *camera.Camera) error {
	return s.streamingSvc.RegisterCamera(ctx, c)
}

func (s *Server) StartInference(ctx context.Context, cameraID, modelID string) (*inference.Job, error) {
	return s.inferenceSvc.StartJob(ctx, cameraID, modelID)
}

func (s *Server) StopInference(ctx context.Context, job *inference.Job) error {
	return s.inferenceSvc.StopJob(ctx, job)
}

func (s *Server) SearchDetections(ctx context.Context, criteria *search.Criteria) ([]*detection.Detection, int64, error) {
	return s.searchSvc.ExecuteSearch(ctx, criteria)
}

func (s *Server) GetModelStatus(ctx context.Context, modelID string) (string, error) {
	return "READY", nil
}
