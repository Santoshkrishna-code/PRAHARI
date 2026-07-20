package grpc

import (
	"context"

	inventoryApp "prahari/services/chemical/internal/application/inventory"
	reportingApp "prahari/services/chemical/internal/application/reporting"
	sdsApp "prahari/services/chemical/internal/application/sds"
	"prahari/services/chemical/internal/domain/chemical"
	"prahari/services/chemical/internal/domain/container"
	sdsDomain "prahari/services/chemical/internal/domain/sds"
)

type Server struct {
	inventorySvc *inventoryApp.Service
	reportingSvc *reportingApp.Service
	sdsSvc       *sdsApp.Service
}

func NewServer(
	inventorySvc *inventoryApp.Service,
	reportingSvc *reportingApp.Service,
	sdsSvc *sdsApp.Service,
) *Server {
	return &Server{
		inventorySvc: inventorySvc,
		reportingSvc: reportingSvc,
		sdsSvc:       sdsSvc,
	}
}

func (s *Server) CreateChemical(ctx context.Context, con *container.Container) error {
	return s.inventorySvc.ReceiveContainer(ctx, con)
}

func (s *Server) ReceiveChemical(ctx context.Context, con *container.Container) error {
	return s.inventorySvc.ReceiveContainer(ctx, con)
}

func (s *Server) IssueChemical(ctx context.Context, containerID, issuedTo string) error {
	return s.inventorySvc.IssueContainer(ctx, containerID, issuedTo)
}

func (s *Server) ReturnChemical(ctx context.Context, containerID string) error {
	return s.inventorySvc.ReturnContainer(ctx, containerID)
}

func (s *Server) GetChemical(ctx context.Context, id string) (*chemical.Chemical, error) {
	return s.reportingSvc.GetChemical(ctx, id)
}

func (s *Server) GetSDS(ctx context.Context, chemicalID string) (*sdsDomain.SDS, error) {
	return s.sdsSvc.GetSDS(ctx, chemicalID)
}
