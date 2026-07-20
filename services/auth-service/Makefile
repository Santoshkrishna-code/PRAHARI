# Auth Service Makefile

.PHONY: run test lint build docker-build docker-dev clean

run:
	go run cmd/server/main.go

test:
	go test -v ./...

lint:
	golangci-lint run ./...

build:
	go build -o bin/auth-service cmd/server/main.go

docker-build:
	docker build -t prahari/auth-service:latest -f Dockerfile .

docker-dev:
	docker build -t prahari/auth-service:dev -f Dockerfile.dev .

clean:
	rm -rf bin/
