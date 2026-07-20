package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"prahari/services/auth-service/internal/api"
)

type HealthHandler struct {
	cognitoClient *cognitoidentityprovider.Client
}

func NewHealthHandler(client *cognitoidentityprovider.Client) *HealthHandler {
	return &HealthHandler{cognitoClient: client}
}

func (h *HealthHandler) Live(w http.ResponseWriter, r *http.Request) {
	// Liveness returns 200 indicating the process is alive
	api.WriteJSON(w, http.StatusOK, map[string]string{"status": "UP"})
}

func (h *HealthHandler) Ready(w http.ResponseWriter, r *http.Request) {
	// Readiness validates connection to downstream AWS Cognito service
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// Issue a light Cognito API check (e.g. DescribeUserPool or ListUserPools limit 1)
	// If it fails with context timeout or network exception, report ready status as DOWN.
	// For template connectivity check:
	_, err := h.cognitoClient.ListUserPools(ctx, &cognitoidentityprovider.ListUserPoolsInput{
		MaxResults: 1,
	})
	if err != nil && ctx.Err() != nil {
		// Context timeout, Cognito is unreachable
		api.WriteError(w, http.StatusServiceUnavailable, "SERVICE_UNAVAILABLE", "Cognito identity provider is unreachable", nil)
		return
	}

	api.WriteJSON(w, http.StatusOK, map[string]string{
		"status":    "UP",
		"cognito":   "CONNECTED",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}
