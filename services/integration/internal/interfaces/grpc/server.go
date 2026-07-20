package grpc

import (
	"context"

	connectorsApp "prahari/services/integration/internal/application/connectors"
	searchApp "prahari/services/integration/internal/application/search"
	synchronizationApp "prahari/services/integration/internal/application/synchronization"
	transformationApp "prahari/services/integration/internal/application/transformation"
	"prahari/services/integration/internal/domain/connector"
	"prahari/services/integration/internal/domain/search"
)

type Server struct {
	connectorsSvc *connectorsApp.Service
	syncSvc       *synchronizationApp.Service
	transformSvc  *transformationApp.Service
	searchSvc     *searchApp.Service
}

func NewServer(
	connectorsSvc *connectorsApp.Service,
	syncSvc *synchronizationApp.Service,
	transformSvc *transformationApp.Service,
	searchSvc *searchApp.Service,
) *Server {
	return &Server{
		connectorsSvc: connectorsSvc,
		syncSvc:       syncSvc,
		transformSvc:  transformSvc,
		searchSvc:     searchSvc,
	}
}

func (s *Server) CreateConnector(ctx context.Context, c *connector.Connector) error {
	return s.connectorsSvc.RegisterConnector(ctx, c)
}

func (s *Server) RunIntegration(ctx context.Context, jobID string) error {
	return s.syncSvc.ExecuteJobSync(ctx, jobID)
}

func (s *Server) TransformMessage(ctx context.Context, connectorID string, raw []byte) ([]byte, error) {
	return s.transformSvc.TransformPayload(ctx, connectorID, raw)
}

func (s *Server) RegisterWebhook(ctx context.Context, targetURL, eventName string) error {
	return nil
}

func (s *Server) SearchIntegrations(ctx context.Context, criteria *search.Criteria) ([]*connector.Connector, int64, error) {
	return s.searchSvc.ExecuteSearch(ctx, criteria)
}
