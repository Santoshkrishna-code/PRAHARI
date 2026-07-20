.PHONY: test lint build run

test:
	go test -v -race -count=1 ./...

lint:
	golangci-lint run ./...

build:
	go build -o bin/server cmd/server/main.go

run:
	go run cmd/server/main.go
