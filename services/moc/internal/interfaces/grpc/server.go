package grpc

import (
	"context"

	reqApp "prahari/services/moc/internal/application/request"
	reportingApp "prahari/services/moc/internal/application/reporting"
	"prahari/services/moc/internal/domain/changerequest"
)

type Server struct {
	reqSvc       *reqApp.Service
	reportingSvc *reportingApp.Service
}

func NewServer(reqSvc *reqApp.Service, reportingSvc *reportingApp.Service) *Server {
	return &Server{
		reqSvc:       reqSvc,
		reportingSvc: reportingSvc,
	}
}

func (s *Server) CreateChangeRequest(ctx context.Context, req *changerequest.Request) error {
	return s.reqSvc.CreateRequest(ctx, req)
}

func (s *Server) GetChangeRequest(ctx context.Context, id string) (*changerequest.Request, error) {
	return s.reportingSvc.GetRequest(ctx, id)
}
