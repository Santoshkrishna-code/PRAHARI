package service_test

import (
	"context"
	"errors"
	"testing"

	"prahari/services/auth-service/internal/domain"
	"prahari/services/auth-service/internal/service"
)

// MockUserRepository implements domain.UserRepository for unit tests.
type MockUserRepository struct {
	SignUpFunc         func(ctx context.Context, email, password, role, firstName, lastName string) (string, error)
	SignInFunc         func(ctx context.Context, email, password string) (*domain.TokenPair, error)
	RefreshTokenFunc   func(ctx context.Context, refreshToken string) (*domain.TokenPair, error)
	GetUserByTokenFunc func(ctx context.Context, accessToken string) (*domain.User, error)
}

func (m *MockUserRepository) SignUp(ctx context.Context, email, password, role, firstName, lastName string) (string, error) {
	return m.SignUpFunc(ctx, email, password, role, firstName, lastName)
}

func (m *MockUserRepository) SignIn(ctx context.Context, email, password string) (*domain.TokenPair, error) {
	return m.SignInFunc(ctx, email, password)
}

func (m *MockUserRepository) RefreshToken(ctx context.Context, refreshToken string) (*domain.TokenPair, error) {
	return m.RefreshTokenFunc(ctx, refreshToken)
}

func (m *MockUserRepository) GetUserByToken(ctx context.Context, accessToken string) (*domain.User, error) {
	return m.GetUserByTokenFunc(ctx, accessToken)
}

// MockUserCache implements repository.UserCache for unit tests.
type MockUserCache struct {
	GetFunc func(ctx context.Context, tokenStr string) (*domain.User, error)
	SetFunc func(ctx context.Context, tokenStr string, user *domain.User) error
}

func (m *MockUserCache) Get(ctx context.Context, tokenStr string) (*domain.User, error) {
	return m.GetFunc(ctx, tokenStr)
}

func (m *MockUserCache) Set(ctx context.Context, tokenStr string, user *domain.User) error {
	return m.SetFunc(ctx, tokenStr, user)
}

func (m *MockUserCache) Delete(ctx context.Context, tokenStr string) error {
	return nil
}

// MockEventPublisher implements events.EventPublisher for unit tests.
type MockEventPublisher struct {
	PublishFunc func(ctx context.Context, topic string, payload interface{}) error
}

func (m *MockEventPublisher) Publish(ctx context.Context, topic string, payload interface{}) error {
	return m.PublishFunc(ctx, topic, payload)
}

func TestRegister_Success(t *testing.T) {
	mockRepo := &MockUserRepository{
		SignUpFunc: func(ctx context.Context, email, password, role, firstName, lastName string) (string, error) {
			return "user-sub-123", nil
		},
	}
	mockCache := &MockUserCache{}
	mockPublisher := &MockEventPublisher{
		PublishFunc: func(ctx context.Context, topic string, payload interface{}) error {
			return nil
		},
	}

	usecase := service.NewAuthUseCase(mockRepo, mockCache, mockPublisher, "dev")
	user, err := usecase.Register(context.Background(), "worker@prahari.internal", "SecurePass1!", "Worker", "John", "Doe")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.ID != "user-sub-123" {
		t.Errorf("expected user ID 'user-sub-123', got %s", user.ID)
	}

	if user.Role != domain.RoleWorker {
		t.Errorf("expected user role 'Worker', got %s", user.Role)
	}
}

func TestRegister_ValidationError(t *testing.T) {
	mockRepo := &MockUserRepository{}
	mockCache := &MockUserCache{}
	mockPublisher := &MockEventPublisher{}

	usecase := service.NewAuthUseCase(mockRepo, mockCache, mockPublisher, "dev")
	_, err := usecase.Register(context.Background(), "", "short", "Worker", "John", "Doe")

	if !errors.Is(err, domain.ErrValidationError) {
		t.Errorf("expected error %v, got %v", domain.ErrValidationError, err)
	}
}

func TestLogin_Success(t *testing.T) {
	mockRepo := &MockUserRepository{
		SignInFunc: func(ctx context.Context, email, password string) (*domain.TokenPair, error) {
			return &domain.TokenPair{
				AccessToken:  "mock-access",
				IDToken:      "mock-id",
				RefreshToken: "mock-refresh",
				ExpiresIn:    3600,
			}, nil
		},
	}
	mockCache := &MockUserCache{}
	mockPublisher := &MockEventPublisher{
		PublishFunc: func(ctx context.Context, topic string, payload interface{}) error {
			return nil
		},
	}

	usecase := service.NewAuthUseCase(mockRepo, mockCache, mockPublisher, "dev")
	tokens, err := usecase.Login(context.Background(), "worker@prahari.internal", "SecurePass1!", "127.0.0.1", "Mozilla/5.0")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if tokens.AccessToken != "mock-access" {
		t.Errorf("expected access token 'mock-access', got %s", tokens.AccessToken)
	}
}
