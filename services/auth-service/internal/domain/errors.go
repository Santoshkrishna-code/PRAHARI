package domain

import "errors"

var (
	ErrInvalidCredentials = errors.New("authentication failed: invalid username or password")
	ErrUserAlreadyExists  = errors.New("registration failed: a user with this email already exists")
	ErrUserNotFound       = errors.New("user entity not found")
	ErrTokenExpired       = errors.New("authorization failed: token is expired")
	ErrInvalidToken       = errors.New("authorization failed: token signature is invalid")
	ErrValidationError    = errors.New("validation failed: invalid input parameters")
	ErrInternalServer     = errors.New("internal server error: processing failed")
)
