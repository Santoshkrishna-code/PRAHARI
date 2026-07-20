package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GRPCStatus maps standard ErrorCode parameters to gRPC code constants.
func GRPCStatus(code ErrorCode) codes.Code {
	switch code {
	case CodeInvalidArgument:
		return codes.InvalidArgument
	case CodeUnauthenticated:
		return codes.Unauthenticated
	case CodePermissionDenied:
		return codes.PermissionDenied
	case CodeNotFound:
		return codes.NotFound
	case CodeAlreadyExists:
		return codes.AlreadyExists
	case CodeConflict:
		return codes.FailedPrecondition
	case CodeResourceExhausted:
		return codes.ResourceExhausted
	case CodeUnavailable:
		return codes.Unavailable
	case CodeDeadlineExceeded:
		return codes.DeadlineExceeded
	case CodeInternal:
		return codes.Internal
	default:
		return codes.Unknown
	}
}

// ToGRPCError translates any native error to a gRPC status.Error.
func ToGRPCError(err error) error {
	if err == nil {
		return nil
	}
	
	// If it is already a gRPC status error, pass it through directly
	if _, ok := status.FromError(err); ok {
		return err
	}
	
	code := GetCode(err)
	grpcCode := GRPCStatus(code)
	
	return status.Error(grpcCode, err.Error())
}
