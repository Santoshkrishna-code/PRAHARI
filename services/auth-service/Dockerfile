# ==============================================================================
# Build Stage
# ==============================================================================
FROM golang:1.21-alpine AS builder

RUN apk add --no-cache git gcc libc-dev

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download || true

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o prahari-auth-service ./cmd/server/main.go

# ==============================================================================
# Run Stage
# ==============================================================================
FROM alpine:3.18

RUN apk --no-cache add ca-certificates tzdata

# Create application user and group
RUN addgroup -S praharigroup && adduser -S prahariuser -G praharigroup
USER prahariuser

WORKDIR /home/prahariuser

COPY --from=builder /app/prahari-auth-service .

ENV PORT=8001
EXPOSE 8001 9001

ENTRYPOINT ["./prahari-auth-service"]
