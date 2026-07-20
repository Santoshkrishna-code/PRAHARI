package grpc

import (
	"context"

	approvalApp "prahari/services/document/internal/application/approval"
	creationApp "prahari/services/document/internal/application/creation"
	lifecycleApp "prahari/services/document/internal/application/lifecycle"
	reportingApp "prahari/services/document/internal/application/reporting"
	versioningApp "prahari/services/document/internal/application/versioning"
	"prahari/services/document/internal/domain/document"
)


type Server struct {
	creationSvc   *creationApp.Service
	lifecycleSvc  *lifecycleApp.Service
	versioningSvc *versioningApp.Service
	approvalSvc   *approvalApp.Service
	reportingSvc  *reportingApp.Service
}

func NewServer(
	creationSvc *creationApp.Service,
	lifecycleSvc *lifecycleApp.Service,
	versioningSvc *versioningApp.Service,
	approvalSvc *approvalApp.Service,
	reportingSvc *reportingApp.Service,
) *Server {
	return &Server{
		creationSvc:   creationSvc,
		lifecycleSvc:  lifecycleSvc,
		versioningSvc: versioningSvc,
		approvalSvc:   approvalSvc,
		reportingSvc:  reportingSvc,
	}
}

func (s *Server) CreateDocument(ctx context.Context, doc *document.Document) error {
	return s.creationSvc.CreateDocument(ctx, doc)
}

func (s *Server) GetDocument(ctx context.Context, id string) (*document.Document, error) {
	return s.reportingSvc.GetDocument(ctx, id)
}

func (s *Server) ApproveDocument(ctx context.Context, id, approverID, comments string) error {
	return s.approvalSvc.ApproveDocument(ctx, id, approverID, comments)
}

func (s *Server) PublishDocument(ctx context.Context, id string) error {
	return s.lifecycleSvc.PublishDocument(ctx, id)
}

func (s *Server) CheckoutDocument(ctx context.Context, id, userID string) error {
	return s.versioningSvc.Checkout(ctx, id, userID)
}
