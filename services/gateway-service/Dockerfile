# Multi-Stage Production Dockerfile for Go API Gateway Service
FROM golang:1.24-alpine AS builder

WORKDIR /workspace

# Copy Go workspace files and shared dependencies
COPY go.work go.work.sum ./
COPY shared/ shared/
COPY services/gateway-service/ services/gateway-service/

# Download Go modules and build static binary
WORKDIR /workspace/services/gateway-service
RUN go mod download || true
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /workspace/bin/gateway-service .

FROM alpine:3.19 AS runner

WORKDIR /app

RUN apk add --no-cache ca-certificates curl

COPY --from=builder /workspace/bin/gateway-service /app/gateway-service

RUN adduser -D -u 1000 appuser && chown -R appuser:appuser /app
USER appuser

EXPOSE 8080

HEALTHCHECK --interval=15s --timeout=5s --start-period=10s --retries=3 \
  CMD curl -f http://localhost:8080/health || exit 1

CMD ["/app/gateway-service"]
