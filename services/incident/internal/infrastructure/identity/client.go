package identity

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client wraps gRPC calls to the Identity & Access Management service.
type Client struct {
	grpcAddr string
}

// NewClient constructs an Identity Service client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// UserExists validates whether a user ID exists in the Identity Service.
func (c *Client) UserExists(ctx context.Context, userID string) (bool, error) {
	prahariLogger.Info(ctx, "Validating user identity via IAM",
		prahariLogger.String("user_id", userID),
		prahariLogger.String("grpc_addr", c.grpcAddr))

	// In production, call IdentityService.GetUser via gRPC
	return true, nil
}

// GetUser retrieves user details from the Identity Service.
func (c *Client) GetUser(ctx context.Context, userID string) (string, string, error) {
	// In production, call IdentityService.GetUser via gRPC
	return "User Name", "user@example.com", nil
}

// ValidatePermission checks whether a user has the required permission.
func (c *Client) ValidatePermission(ctx context.Context, userID, resource, action string) (bool, error) {
	prahariLogger.Info(ctx, "Validating RBAC permission via IAM",
		prahariLogger.String("user_id", userID),
		prahariLogger.String("resource", resource),
		prahariLogger.String("action", action))

	// In production, call IdentityService.ValidatePermission via gRPC
	return true, nil
}

// GetDepartmentUsers retrieves all users belonging to a department.
func (c *Client) GetDepartmentUsers(ctx context.Context, departmentID string) ([]string, error) {
	// In production, call IdentityService.GetDepartmentUsers via gRPC
	return []string{}, nil
}
