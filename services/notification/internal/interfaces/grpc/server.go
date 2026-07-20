package grpc

import (
	"context"
	"errors"

	prahariFlow "prahari/services/notification/internal/application/dispatcher"
)

// Server implements gRPC messaging handlers.
type Server struct {
	flow *prahariFlow.Flow
}

// NewServer constructs a Server.
func NewServer(flow *prahariFlow.Flow) *Server {
	return &Server{flow: flow}
}

// SendNotification dispatches messages via gRPC calls.
func (s *Server) SendNotification(ctx context.Context, channelName, recipient, body string) error {
	if recipient == "" {
		return errors.New("recipient coordinate is required")
	}

	return s.flow.Dispatch(ctx, channelName, recipient, body)
}
