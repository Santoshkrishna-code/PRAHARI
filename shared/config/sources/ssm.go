package sources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// SSMAPI defines the narrow AWS SSM Client capability needed to fetch parameters.
type SSMAPI interface {
	GetParameter(ctx context.Context, params *ssm.GetParameterInput, optFns ...func(*ssm.Options)) (*ssm.GetParameterOutput, error)
}

// LoadFromSSM retrieves a single string value from the AWS SSM Parameter Store.
func LoadFromSSM(ctx context.Context, client SSMAPI, parameterName string) (string, error) {
	if client == nil {
		return "", fmt.Errorf("ssm client is uninitialized")
	}

	input := &ssm.GetParameterInput{
		Name:           &parameterName,
		WithDecryption: ssm.GetParameterInput{}.WithDecryption, // Defaults to decrypted retrieval for secure parameters
	}
	// Explicitly assign dereferenced boolean
	decrypt := true
	input.WithDecryption = &decrypt

	output, err := client.GetParameter(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to fetch SSM parameter %s: %w", parameterName, err)
	}

	if output.Parameter == nil || output.Parameter.Value == nil {
		return "", fmt.Errorf("parameter %s returned empty value", parameterName)
	}

	return *output.Parameter.Value, nil
}
