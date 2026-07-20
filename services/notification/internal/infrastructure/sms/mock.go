package sms

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Adapter implements SMS API client logic.
type Adapter struct {
}

// NewAdapter constructs an Adapter.
func NewAdapter() *Adapter {
	return &Adapter{}
}

// Send simulates SMS loops.
func (a *Adapter) Send(ctx context.Context, recipient, body string) error {
	prahariLogger.Info(ctx, "Sending SMS message to recipient",
		prahariLogger.String("recipient", recipient),
		prahariLogger.String("body", body))

	return nil
}

// Name returns adapter name tag.
func (a *Adapter) Name() string {
	return "sms"
}
