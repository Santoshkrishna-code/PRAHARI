package workflow

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client wraps gRPC calls to the Workflow Engine service.
type Client struct {
	grpcAddr string
}

// NewClient constructs a Workflow Engine client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// StartWorkflow triggers a workflow instance in the Workflow Engine for an incident.
func (c *Client) StartWorkflow(ctx context.Context, incidentID, workflowType string) error {
	prahariLogger.Info(ctx, "Triggering workflow in Workflow Engine",
		prahariLogger.String("incident_id", incidentID),
		prahariLogger.String("workflow_type", workflowType),
		prahariLogger.String("grpc_addr", c.grpcAddr))

	// In production, establish gRPC connection and call WorkflowService.StartWorkflow:
	// conn, _ := grpc.Dial(c.grpcAddr, grpc.WithInsecure())
	// defer conn.Close()
	// client := pb.NewWorkflowServiceClient(conn)
	// _, err := client.StartWorkflow(ctx, &pb.StartWorkflowRequest{
	//     ReferenceID: incidentID,
	//     WorkflowType: workflowType,
	// })
	return nil
}
