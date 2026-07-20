package dynamodb_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	prahariDynamo "prahari/shared/aws/dynamodb"
)

func TestClient_Ping_Uninitialized(t *testing.T) {
	client := prahariDynamo.NewClient(nil)
	err := client.Ping(context.Background())
	if err == nil {
		t.Error("expected error checking health on nil DynamoDB client, got nil")
	}
}

func TestClient_CRUD_Uninitialized(t *testing.T) {
	client := prahariDynamo.NewClient(nil)
	ctx := context.Background()

	type Item struct {
		ID string `dynamodbav:"id"`
	}

	err := client.PutItem(ctx, "table", Item{ID: "1"})
	if err == nil {
		t.Error("expected put to error on uninitialized client, got nil")
	}

	key := map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{Value: "1"},
	}

	_, err = client.GetItem(ctx, "table", key, &Item{})
	if err == nil {
		t.Error("expected get to error on uninitialized client, got nil")
	}

	err = client.DeleteItem(ctx, "table", key)
	if err == nil {
		t.Error("expected delete to error on uninitialized client, got nil")
	}
}
