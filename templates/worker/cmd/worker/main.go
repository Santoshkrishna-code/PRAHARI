package main

import (
	"context"
	"fmt"

	"prahari/templates/worker/internal/polling"
)

func main() {
	ctx := context.Background()
	fmt.Println("Bootstrapping PRAHARI background worker...")
	_ = polling.PollEventsLoop(ctx)
}
