package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"

	"prahari/services/auth-service/internal/domain"
)

// CognitoUserRepository implements domain.UserRepository interface using the AWS SDK v2.
type CognitoUserRepository struct {
	client     *cognitoidentityprovider.Client
	userPoolID string
	clientID   string
}

func NewCognitoUserRepository(
	client *cognitoidentityprovider.Client,
	userPoolID, clientID string,
) domain.UserRepository {
	return &CognitoUserRepository{
		client:     client,
		userPoolID: userPoolID,
		clientID:   clientID,
	}
}

func (r *CognitoUserRepository) SignUp(
	ctx context.Context,
	email, password, role, firstName, lastName string,
) (string, error) {
	input := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(r.clientID),
		Username: aws.String(email), // Cognito uses email as username in our pool setup
		Password: aws.String(password),
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(email)},
			{Name: aws.String("custom:role"), Value: aws.String(role)},
			{Name: aws.String("given_name"), Value: aws.String(firstName)},
			{Name: aws.String("family_name"), Value: aws.String(lastName)},
		},
	}

	output, err := r.client.SignUp(ctx, input)
	if err != nil {
		var usernameExistsErr *types.UsernameExistsException
		if errors.As(err, &usernameExistsErr) {
			return "", domain.ErrUserAlreadyExists
		}
		var invalidPasswordErr *types.InvalidPasswordException
		if errors.As(err, &invalidPasswordErr) {
			return "", fmt.Errorf("%w: password does not meet complexity requirements", domain.ErrValidationError)
		}
		return "", fmt.Errorf("cognito signup error: %w", err)
	}

	return aws.ToString(output.UserSub), nil
}

func (r *CognitoUserRepository) SignIn(
	ctx context.Context,
	email, password string,
) (*domain.TokenPair, error) {
	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		ClientId: aws.String(r.clientID),
		AuthParameters: map[string]string{
			"USERNAME": email,
			"PASSWORD": password,
		},
	}

	output, err := r.client.InitiateAuth(ctx, input)
	if err != nil {
		var notAuthorizedErr *types.NotAuthorizedException
		var userNotFoundErr *types.UserNotFoundException
		if errors.As(err, &notAuthorizedErr) || errors.As(err, &userNotFoundErr) {
			return nil, domain.ErrInvalidCredentials
		}
		return nil, fmt.Errorf("cognito signin error: %w", err)
	}

	if output.AuthenticationResult == nil {
		return nil, fmt.Errorf("authentication incomplete: MFA or challenges required")
	}

	res := output.AuthenticationResult
	return &domain.TokenPair{
		AccessToken:  aws.ToString(res.AccessToken),
		IDToken:      aws.ToString(res.IdToken),
		RefreshToken: aws.ToString(res.RefreshToken),
		ExpiresIn:    res.ExpiresIn,
	}, nil
}

func (r *CognitoUserRepository) RefreshToken(
	ctx context.Context,
	refreshToken string,
) (*domain.TokenPair, error) {
	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeRefreshTokenAuth,
		ClientId: aws.String(r.clientID),
		AuthParameters: map[string]string{
			"REFRESH_TOKEN": refreshToken,
		},
	}

	output, err := r.client.InitiateAuth(ctx, input)
	if err != nil {
		var notAuthorizedErr *types.NotAuthorizedException
		if errors.As(err, &notAuthorizedErr) {
			return nil, domain.ErrInvalidToken
		}
		return nil, fmt.Errorf("cognito token refresh error: %w", err)
	}

	res := output.AuthenticationResult
	if res == nil {
		return nil, fmt.Errorf("failed to retrieve refresh result")
	}

	return &domain.TokenPair{
		AccessToken:  aws.ToString(res.AccessToken),
		IDToken:      aws.ToString(res.IdToken),
		RefreshToken: refreshToken, // Refresh flow doesn't always return a new refresh token, reuse current
		ExpiresIn:    res.ExpiresIn,
	}, nil
}

func (r *CognitoUserRepository) GetUserByToken(
	ctx context.Context,
	accessToken string,
) (*domain.User, error) {
	input := &cognitoidentityprovider.GetUserInput{
		AccessToken: aws.String(accessToken),
	}

	output, err := r.client.GetUser(ctx, input)
	if err != nil {
		var notAuthorizedErr *types.NotAuthorizedException
		if errors.As(err, &notAuthorizedErr) {
			return nil, domain.ErrInvalidToken
		}
		return nil, fmt.Errorf("cognito getuser error: %w", err)
	}

	user := &domain.User{
		ID:    aws.ToString(output.Username),
		Email: aws.ToString(output.Username), // Email is primary username
	}

	for _, attr := range output.UserAttributes {
		name := aws.ToString(attr.Name)
		val := aws.ToString(attr.Value)
		
		switch name {
		case "sub":
			user.ID = val
		case "email":
			user.Email = val
		case "custom:role":
			user.Role = domain.UserRole(val)
		case "given_name":
			user.FirstName = val
		case "family_name":
			user.LastName = val
		}
	}

	return user, nil
}
