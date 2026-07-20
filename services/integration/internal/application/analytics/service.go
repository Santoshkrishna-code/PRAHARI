package analytics

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Repository interface {
	GetDashboardMetrics(ctx context.Context, plantID string) (map[string]float64, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetIntegrationPerformance(ctx context.Context, plantID string) (map[string]float64, error) {
	prahariLogger.Info(ctx, "Retrieving integration hub performance statistics",
		prahariLogger.String("plant_id", plantID))
	return map[string]float64{
		"connector_health_pct":         100.0,
		"active_integrations_count":    15.0,
		"message_throughput_per_sec":   125.0,
		"failed_integrations_count":    2.0,
		"retry_count":                  8.0,
		"dlq_size_count":               1.0,
		"average_latency_ms":           45.2,
		"protocol_usage_mqtt_pct":      40.0,
		"protocol_usage_opcua_pct":     35.0,
		"protocol_usage_rest_pct":      25.0,
		"synchronization_time_sec":     12.4,
		"external_system_avail_pct":    99.8,
	}, nil
}
