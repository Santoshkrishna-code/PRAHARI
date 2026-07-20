package cloudwatch_test

import (
	"context"
	"testing"

	prahariCW "prahari/shared/aws/cloudwatch"
)

func TestClient_Ping_Uninitialized(t *testing.T) {
	client := prahariCW.NewClient(nil)
	err := client.Ping(context.Background())
	if err == nil {
		t.Error("expected error checking health on nil CloudWatch client, got nil")
	}
}

func TestClient_PutMetric_Uninitialized(t *testing.T) {
	client := prahariCW.NewClient(nil)
	err := client.PutMetricData(context.Background(), "ns", "metric", 10.5, "Seconds")
	if err == nil {
		t.Error("expected metric push to error on uninitialized client, got nil")
	}
}
