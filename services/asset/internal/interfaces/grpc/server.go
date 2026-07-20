package grpc

import (
	"context"
	"errors"

	assetApp "prahari/services/asset/internal/application/asset"
	searchApp "prahari/services/asset/internal/application/search"
	searchDomain "prahari/services/asset/internal/domain/search"
	assetDomain "prahari/services/asset/internal/domain/asset"
)

// Server exposes gRPC endpoints.
type Server struct {
	asset  *assetApp.Service
	search *searchApp.Service
}

// NewServer instantiates Server.
func NewServer(
	asset *assetApp.Service,
	search *searchApp.Service,
) *Server {
	return &Server{
		asset:  asset,
		search: search,
	}
}

// CreateAsset inserts profile.
func (s *Server) CreateAsset(ctx context.Context, cmd assetApp.RegisterAssetCommand) (*assetDomain.Asset, error) {
	return s.asset.RegisterAsset(ctx, cmd, "grpc-actor")
}

// GetAsset returns asset profile details.
func (s *Server) GetAsset(ctx context.Context, id string) (*assetDomain.Asset, error) {
	if id == "" {
		return nil, errors.New("asset ID is required")
	}
	return s.asset.GetAsset(ctx, id)
}

// VerifyAsset check lifecycle permit state validations.
func (s *Server) VerifyAsset(ctx context.Context, id string) (bool, error) {
	a, err := s.asset.GetAsset(ctx, id)
	if err != nil {
		return false, err
	}
	return a.LifecycleStatus == "OPERATIONAL", nil
}

// ChangeLifecycle status.
func (s *Server) ChangeLifecycle(ctx context.Context, id, targetCode, actor string) error {
	cmd := assetApp.TransitionStatusCommand{
		AssetID:    id,
		TargetCode: targetCode,
		ActorID:    actor,
	}
	return s.asset.TransitionLifecycle(ctx, cmd)
}

// SearchAssets query matches.
func (s *Server) SearchAssets(ctx context.Context, criteria *searchDomain.Criteria) (*searchDomain.Result, error) {
	return s.search.Search(ctx, criteria)
}
