# Distributed Caching Subsystem

This package implements distributed read-through/write-through cache wrappers supporting JSON/GOB encoding, eviction keys, TTL trackers, and pipeline batch operations.

---

## 1. Features

- **Decoupled Serializer**: Swaps between raw JSON or GOB binaries.
- **Cache Eviction**: evicts keys.
- **TTL Metrics**: Tracks remaining durations.
- **Pipeline Batching**: Optimizes network trips using pipelined MGet and MSet.

---

## 2. API Examples

### A. Initializing and accessing caches
```go
import (
	"prahari/shared/redis"
	"prahari/shared/redis/cache"
	"prahari/shared/redis/serializer"
)

client, _ := redis.NewClient(redisConfig)
ser := serializer.NewJSONSerializer()
c := cache.NewCache(client, ser)

// Write cache item
_ = c.Set(ctx, "user:123", userStruct, 15*time.Minute)

// Read cache item
var target User
err := c.Get(ctx, "user:123", &target)
if err != nil {
    if errors.Is(err, cache.ErrCacheMiss) {
        // Cache miss
    }
}
```
