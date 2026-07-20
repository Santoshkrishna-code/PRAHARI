.PHONY: test lint build run

test:
	GOWORK=off go test -v -race -count=1 ./...

lint:
	golangci-lint run ./...

build:
	GOWORK=off go build -o bin/server cmd/server/main.go

run:
	GOWORK=off go run cmd/server/main.go
