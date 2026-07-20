# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY shared/ /shared
COPY services/observation/ /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /server cmd/server/main.go

# Production stage
FROM alpine:3.18
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /
COPY --from=builder /server /server
EXPOSE 8093
ENTRYPOINT ["/server"]
