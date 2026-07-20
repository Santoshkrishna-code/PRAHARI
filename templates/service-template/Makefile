# Service-specific Makefile

.PHONY: test lint build docker-build run

test:
	go test -v ./...

lint:
	golangci-lint run ./...

build:
	go build -o bin/service cmd/main.go

docker-build:
	docker build -t prahari/service-template:latest .

run:
	go run cmd/main.go
