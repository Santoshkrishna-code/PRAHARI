package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Client wraps the DynamoDB SDK client.
type Client struct {
	client *dynamodb.Client
}

// NewClient constructs a DynamoDB client wrapper.
func NewClient(client *dynamodb.Client) *Client {
	return &Client{client: client}
}

// Ping implements aws.HealthChecker checking connectivity by listing tables.
func (c *Client) Ping(ctx context.Context) error {
	if c.client == nil {
		return fmt.Errorf("dynamodb client is uninitialized")
	}
	_, err := c.client.ListTables(ctx, &dynamodb.ListTablesInput{Limit: &[]int32{1}[0]})
	return err
}

// PutItem marshals and inserts/replaces a record in the specified DynamoDB table.
func (c *Client) PutItem(ctx context.Context, tableName string, item interface{}) error {
	if c.client == nil {
		return fmt.Errorf("dynamodb client is uninitialized")
	}

	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("failed to marshal item: %w", err)
	}

	_, err = c.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &tableName,
		Item:      av,
	})
	if err != nil {
		return fmt.Errorf("failed to put item to %s: %w", tableName, err)
	}

	return nil
}

// GetItem retrieves a record matching the primary key, unmarshaling it into the target destination.
// Returns boolean indicating if the item was found.
func (c *Client) GetItem(ctx context.Context, tableName string, key map[string]types.AttributeValue, target interface{}) (bool, error) {
	if c.client == nil {
		return false, fmt.Errorf("dynamodb client is uninitialized")
	}

	output, err := c.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &tableName,
		Key:       key,
	})
	if err != nil {
		return false, fmt.Errorf("failed to get item from %s: %w", tableName, err)
	}

	if output.Item == nil {
		return false, nil // Not found
	}

	err = attributevalue.UnmarshalMap(output.Item, target)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal item: %w", err)
	}

	return true, nil
}

// DeleteItem removes a record matching the primary key.
func (c *Client) DeleteItem(ctx context.Context, tableName string, key map[string]types.AttributeValue) error {
	if c.client == nil {
		return fmt.Errorf("dynamodb client is uninitialized")
	}

	_, err := c.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &tableName,
		Key:       key,
	})
	if err != nil {
		return fmt.Errorf("failed to delete item from %s: %w", tableName, err)
	}

	return nil
}

// TransactWriteItems executes multiple DynamoDB write actions (Put/Update/Delete) atomically.
func (c *Client) TransactWriteItems(ctx context.Context, actions []types.TransactWriteItem) error {
	if c.client == nil {
		return fmt.Errorf("dynamodb client is uninitialized")
	}

	_, err := c.client.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: actions,
	})
	if err != nil {
		return fmt.Errorf("dynamodb transaction failed: %w", err)
	}

	return nil
}
