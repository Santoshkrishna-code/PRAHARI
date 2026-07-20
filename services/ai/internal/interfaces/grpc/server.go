package grpc

import (
	"context"

	chatApp "prahari/services/ai/internal/application/chat"
	predictionApp "prahari/services/ai/internal/application/prediction"
	ragApp "prahari/services/ai/internal/application/rag"
	recommendationApp "prahari/services/ai/internal/application/recommendation"
	searchApp "prahari/services/ai/internal/application/search"
	summarizationApp "prahari/services/ai/internal/application/summarization"
	"prahari/services/ai/internal/domain/conversation"
	"prahari/services/ai/internal/domain/document"
	"prahari/services/ai/internal/domain/generation"
	"prahari/services/ai/internal/domain/prediction"
	"prahari/services/ai/internal/domain/recommendation"
	"prahari/services/ai/internal/domain/retrieval"
	"prahari/services/ai/internal/domain/search"
	"prahari/services/ai/internal/domain/summarization"
)

type Server struct {
	chatSvc           *chatApp.Service
	ragSvc            *ragApp.Service
	summarizationSvc  *summarizationApp.Service
	recommendationSvc *recommendationApp.Service
	predictionSvc     *predictionApp.Service
	searchSvc         *searchApp.Service
}

func NewServer(
	chatSvc *chatApp.Service,
	ragSvc *ragApp.Service,
	summarizationSvc *summarizationApp.Service,
	recommendationSvc *recommendationApp.Service,
	predictionSvc *predictionApp.Service,
	searchSvc *searchApp.Service,
) *Server {
	return &Server{
		chatSvc:           chatSvc,
		ragSvc:            ragSvc,
		summarizationSvc:  summarizationSvc,
		recommendationSvc: recommendationSvc,
		predictionSvc:     predictionSvc,
		searchSvc:         searchSvc,
	}
}

func (s *Server) Chat(ctx context.Context, threadID, role, content string) (*generation.Response, error) {
	return s.chatSvc.SubmitMessage(ctx, threadID, role, content)
}

func (s *Server) RetrieveContext(ctx context.Context, query string, limit int) ([]*retrieval.Result, error) {
	return s.ragSvc.SearchContext(ctx, query, limit)
}

func (s *Server) GenerateSummary(ctx context.Context, sourceID, original string) (*summarization.Summary, error) {
	return s.summarizationSvc.SummarizeDocument(ctx, sourceID, original)
}

func (s *Server) GenerateRecommendation(ctx context.Context, plantID, sourceID, recType string) (*recommendation.Recommendation, error) {
	return s.recommendationSvc.GenerateSuggestions(ctx, plantID, sourceID, recType)
}

func (s *Server) Predict(ctx context.Context, plantID, topic string) (*prediction.Forecast, error) {
	return s.predictionSvc.PredictRisk(ctx, plantID, topic)
}

func (s *Server) SearchKnowledge(ctx context.Context, criteria *search.Criteria) ([]*document.Doc, int64, error) {
	return s.searchSvc.ExecuteSearch(ctx, criteria)
}

func (s *Server) CreateThread(ctx context.Context, userID, plantID string) (*conversation.Thread, error) {
	return s.chatSvc.CreateThread(ctx, userID, plantID)
}
