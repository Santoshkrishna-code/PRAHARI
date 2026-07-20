# ==============================================================================
# Stage 1: Build binary statically
# ==============================================================================
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git gcc libc-dev

WORKDIR /app

# Cache module layers
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Run validation checks and compile statically
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o prahari-service ./cmd/main.go

# ==============================================================================
# Stage 2: Minimal runtime image
# ==============================================================================
FROM alpine:3.18

RUN apk --no-cache add ca-certificates tzdata

# Run as non-root user for container security hardening
RUN addgroup -S praharigroup && adduser -S prahariuser -G praharigroup
USER prahariuser

WORKDIR /home/prahariuser

COPY --from=builder /app/prahari-service .

ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["./prahari-service"]
