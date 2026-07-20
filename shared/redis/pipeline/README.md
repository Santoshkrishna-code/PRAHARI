# Redis Pipelines & Transactions Subsystem

This package wraps Redis Pipeline commands and MULTI/EXEC optimistic transaction loops.

---

## 1. Features

- **Pipeliner**: Queues commands, saving network bandwidth.
- **Optimistic Locking**: Monitors keys for concurrent edits using WATCH.

---

## 2. API Examples

```go
import "prahari/shared/redis/pipeline"

pipe, _ := pipeline.NewPipeline(client)
pipe.Set(ctx, "a", "1", 0)
pipe.Set(ctx, "b", "2", 0)
_, _ = pipe.Exec(ctx)
```
