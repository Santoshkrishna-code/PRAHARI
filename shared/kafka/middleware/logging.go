package middleware

import (
	"context"
	"fmt"
	"time"
)

// LoggingMiddleware wraps a message handler to log processing latency, keys, and execution status.
func LoggingMiddleware(handler func(ctx context.Context, key, val []byte) error) func(context.Context, key, val []byte) error {
	return func(ctx context.Context, key, val []byte) error {
		start := time.Now()
		
		fmt.Printf("[KAFKA-MID] Starting message processing. Key: %s\n", string(key))
		
		err := handler(ctx, key, val)
		
		elapsed := time.Since(start)
		if err != nil {
			fmt.Printf("[KAFKA-MID] Message processing failed. Key: %s, Duration: %v, Error: %v\n", string(key), elapsed, err)
		} else {
			fmt.Printf("[KAFKA-MID] Message processing succeeded. Key: %s, Duration: %v\n", string(key), elapsed)
		}

		return err
	}
}
