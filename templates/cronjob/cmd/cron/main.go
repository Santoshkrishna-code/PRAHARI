package main

import (
	"context"
	"fmt"

	"prahari/templates/cronjob/internal/task"
)

func main() {
	fmt.Println("Running PRAHARI scheduled task...")
	_ = task.RunTask(context.Background())
}
