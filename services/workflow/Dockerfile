# Build Stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy shared sdk library dependency
COPY shared/ /shared

# Copy service codes
COPY services/workflow/ /app

# Build clean static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /server cmd/server/main.go

# Runtime Stage
FROM alpine:3.18

WORKDIR /

COPY --from=builder /server /server

# Expose server port
EXPOSE 8082

ENTRYPOINT ["/server"]
