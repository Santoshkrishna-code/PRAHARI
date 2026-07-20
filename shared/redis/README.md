# PRAHARI Shared Redis Platform SDK

This package provides a connection pool, distributed caching client, distributed Redlocks, channel publishers/subscribers, rate limiters, serializers, and pipelined transactions.

---

## 1. Directory Structure

- **`cache/`**: GET, SET, DELETE, TTL, and pipelined MGet/MSet.
- **`locks/`**: Distributed Mutex, Lua script unlocks, lease watchdogs, and Redlocks.
- **`pubsub/`**: Publish/Subscribe message streams.
- **`ratelimit/`**: Fixed-window, Sliding-window log, and Token Bucket limiters.
- **`serializer/`**: JSON and GOB format drivers.
- **`pipeline/`**: Batched commands and transactions.

---

## 2. Usage Blueprint

Exposes connection checks (`Ping`) and manages pool reconnections.
```go
import (
    "prahari/shared/redis"
    "prahari/shared/redis/cache"
    "prahari/shared/redis/serializer"
)

client, _ := redis.NewClient(redis.Config{Address: "localhost:6379"})
defer client.Close()

// Setup cache client
c := cache.NewCache(client, serializer.NewJSONSerializer())
```
