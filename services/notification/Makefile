.PHONY: test lint build

test:
	go test -v ./...

lint:
	golangci-lint run ./...

build:
	go build -o bin/server cmd/server/main.go
