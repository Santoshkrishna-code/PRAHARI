.PHONY: format lint test tidy security verify

# Default target
verify: format lint test tidy security

format:
	go fmt ./...

lint:
	golangci-lint run ./...

test:
	go test -race -v ./...

tidy:
	go mod tidy

security:
	gosec ./...
