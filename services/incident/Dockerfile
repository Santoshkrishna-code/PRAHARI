# Build Stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy shared SDK dependency
COPY shared/ /shared

# Copy service source
COPY services/incident/ /app

# Build clean static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /server cmd/server/main.go

# Runtime Stage
FROM alpine:3.18

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /

COPY --from=builder /server /server

# Expose service port
EXPOSE 8084

ENTRYPOINT ["/server"]
