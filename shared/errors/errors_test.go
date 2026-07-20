package errors_test

import (
	"database/sql"
	"errors"
	"fmt"
	"net"
	"strings"
	"testing"

	prahariErrors "prahari/shared/errors"
)

func TestNewAppError(t *testing.T) {
	err := prahariErrors.New(prahariErrors.CodeInvalidArgument, "invalid input parameter value")

	if err.Code != prahariErrors.CodeInvalidArgument {
		t.Errorf("expected code %s, got %s", prahariErrors.CodeInvalidArgument, err.Code)
	}

	if err.Message != "invalid input parameter value" {
		t.Errorf("expected message, got %s", err.Message)
	}

	if len(err.Stack) == 0 {
		t.Error("expected stack trace to be captured")
	}
}

func TestWrap(t *testing.T) {
	rootErr := errors.New("raw connection failure")
	wrapped := prahariErrors.Wrap(rootErr, prahariErrors.CodeInternal, "failed connection test")

	if !errors.Is(wrapped, rootErr) {
		t.Error("expected wrapped error to be unpackable with errors.Is")
	}

	if prahariErrors.Cause(wrapped) != rootErr {
		t.Error("expected Cause() helper to locate the root raw error")
	}
}

func TestClassification(t *testing.T) {
	err := prahariErrors.New(prahariErrors.CodeUnavailable, "system busy")
	if prahariErrors.IsRetryable(err) {
		t.Error("expected new error to default to non-retryable status")
	}

	retryErr := prahariErrors.MarkRetryable(err)
	if !prahariErrors.IsRetryable(retryErr) {
		t.Error("expected marked error to return true for IsRetryable check")
	}
}

func TestProblemDetails(t *testing.T) {
	err := prahariErrors.New(prahariErrors.CodeNotFound, "asset missing")
	problem := prahariErrors.NewProblemDetails(err, 404, "Asset Not Found", "/assets/123")

	if problem.Status != 404 {
		t.Errorf("expected HTTP status 404, got %d", problem.Status)
	}

	if problem.Code != string(prahariErrors.CodeNotFound) {
		t.Errorf("expected error code string, got %s", problem.Code)
	}
}

func TestTranslateDatabaseError(t *testing.T) {
	err := prahariErrors.TranslateDatabaseError(sql.ErrNoRows, "Permit")
	if prahariErrors.GetCode(err) != prahariErrors.CodeNotFound {
		t.Errorf("expected database missing error mapped to CodeNotFound, got %s", prahariErrors.GetCode(err))
	}
}

type mockAddr struct{}

func (mockAddr) Network() string { return "tcp" }
func (mockAddr) String() string  { return "127.0.0.1:80" }

type mockNetErr struct {
	timeout bool
}

func (e mockNetErr) Error() string   { return "connection timeout" }
func (e mockNetErr) Timeout() bool   { return e.timeout }
func (e mockNetErr) Temporary() bool { return true }

func TestTranslateNetworkError(t *testing.T) {
	rawNetErr := &net.OpError{
		Op:  "dial",
		Net: "tcp",
		Err: mockNetErr{timeout: true},
	}
	
	err := prahariErrors.TranslateNetworkError(rawNetErr)
	if prahariErrors.GetCode(err) != prahariErrors.CodeUnavailable {
		t.Errorf("expected CodeUnavailable for network failure, got %s", prahariErrors.GetCode(err))
	}
	
	if !prahariErrors.IsRetryable(err) {
		t.Error("expected timeout net.OpError to be classified as retryable")
	}
}

func TestFormatStackTrace(t *testing.T) {
	err := prahariErrors.New(prahariErrors.CodeInternal, "debug crash")
	formatted := fmt.Sprintf("%+v", err)
	
	if !strings.Contains(formatted, "debug crash") {
		t.Error("expected formatted string to contain error message")
	}
	
	if !strings.Contains(formatted, "Stack Trace:") {
		t.Error("expected formatted string %+v to output stack trace title header")
	}
}
