package cloudwatch

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

// Client wraps the AWS CloudWatch SDK client.
type Client struct {
	client *cloudwatch.Client
}

// NewClient constructs a CloudWatch client wrapper.
func NewClient(client *cloudwatch.Client) *Client {
	return &Client{client: client}
}

// Ping implements aws.HealthChecker checking connectivity by describing alarm history.
func (c *Client) Ping(ctx context.Context) error {
	if c.client == nil {
		return fmt.Errorf("cloudwatch client is uninitialized")
	}
	_, err := c.client.DescribeAlarmHistory(ctx, &cloudwatch.DescribeAlarmHistoryInput{MaxRecords: &[]int32{1}[0]})
	return err
}

// PutMetricData posts a single custom metric value into the specified CloudWatch namespace.
func (c *Client) PutMetricData(ctx context.Context, namespace, metricName string, value float64, unit string) error {
	if c.client == nil {
		return fmt.Errorf("cloudwatch client is uninitialized")
	}

	metricUnit := types.StandardUnit(unit)

	datum := types.MetricDatum{
		MetricName: &metricName,
		Value:      &value,
		Unit:       metricUnit,
	}

	input := &cloudwatch.PutMetricDataInput{
		Namespace:  &namespace,
		MetricData: []types.MetricDatum{datum},
	}

	_, err := c.client.PutMetricData(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to put CloudWatch metric data: %w", err)
	}

	return nil
}
