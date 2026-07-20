package kafka

import (
	"context"
	"fmt"
)

// Ping validates TCP connections to the first configured Kafka broker.
func Ping(ctx context.Context, cfg Config) error {
	if len(cfg.Brokers) == 0 {
		return fmt.Errorf("no brokers configured")
	}

	dialer, err := GetDialer(cfg)
	if err != nil {
		return err
	}

	conn, err := dialer.DialContext(ctx, "tcp", cfg.Brokers[0])
	if err != nil {
		return fmt.Errorf("failed to connect to broker %s: %w", cfg.Brokers[0], err)
	}
	defer conn.Close()

	return nil
}
