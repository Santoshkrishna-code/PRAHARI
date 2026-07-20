package main

import (
	"context"
	"fmt"

	"prahari/templates/lambda/internal/handler"
)

func main() {
	fmt.Println("Starting serverless execution runtime...")
	req := handler.Request{EventID: "ev-999"}
	res, _ := handler.HandleRequest(context.Background(), req)
	fmt.Printf("Result: %s\n", res)
}
