package aws

import (
	"context"
)

// CheckAWSHealth iterates over multiple service health pings and returns errors mapping.
func CheckAWSHealth(ctx context.Context, services map[string]HealthChecker) map[string]error {
	results := make(map[string]error)
	for name, service := range services {
		if service == nil {
			continue
		}
		// If ping takes too long or errors out, record in the results map
		if err := service.Ping(ctx); err != nil {
			results[name] = err
		} else {
			results[name] = nil
		}
	}
	return results
}
