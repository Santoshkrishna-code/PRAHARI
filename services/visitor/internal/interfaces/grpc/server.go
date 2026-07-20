package grpc

import (
	"context"

	approvalApp "prahari/services/visitor/internal/application/approval"
	checkinApp "prahari/services/visitor/internal/application/checkin"
	checkoutApp "prahari/services/visitor/internal/application/checkout"
	musterApp "prahari/services/visitor/internal/application/muster"
	registrationApp "prahari/services/visitor/internal/application/registration"
	reportingApp "prahari/services/visitor/internal/application/reporting"
	"prahari/services/visitor/internal/domain/emergencymuster"
	"prahari/services/visitor/internal/domain/visitor"
)

type Server struct {
	registrationSvc *registrationApp.Service
	approvalSvc     *approvalApp.Service
	checkinSvc      *checkinApp.Service
	checkoutSvc     *checkoutApp.Service
	musterSvc       *musterApp.Service
	reportingSvc    *reportingApp.Service
}

func NewServer(
	registrationSvc *registrationApp.Service,
	approvalSvc *approvalApp.Service,
	checkinSvc *checkinApp.Service,
	checkoutSvc *checkoutApp.Service,
	musterSvc *musterApp.Service,
	reportingSvc *reportingApp.Service,
) *Server {
	return &Server{
		registrationSvc: registrationSvc,
		approvalSvc:     approvalSvc,
		checkinSvc:      checkinSvc,
		checkoutSvc:     checkoutSvc,
		musterSvc:       musterSvc,
		reportingSvc:    reportingSvc,
	}
}

func (s *Server) CreateVisitor(ctx context.Context, vis *visitor.Visitor) error {
	return s.registrationSvc.RegisterVisitor(ctx, vis)
}

func (s *Server) ApproveVisitor(ctx context.Context, visitID string) error {
	return s.approvalSvc.ApproveVisit(ctx, visitID)
}

func (s *Server) CheckInVisitor(ctx context.Context, visitID, checkpoint, operatorID string) error {
	return s.checkinSvc.CheckIn(ctx, visitID, checkpoint, operatorID)
}

func (s *Server) CheckOutVisitor(ctx context.Context, visitID, operatorID string) error {
	return s.checkoutSvc.CheckOut(ctx, visitID, operatorID)
}

func (s *Server) GenerateMusterReport(ctx context.Context, rec *emergencymuster.Record) error {
	return s.musterSvc.AccountForVisitor(ctx, rec)
}

func (s *Server) GetVisitor(ctx context.Context, id string) (*visitor.Visitor, error) {
	return s.reportingSvc.GetVisitor(ctx, id)
}
