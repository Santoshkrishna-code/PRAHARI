# Serialization Subsystem

This package provides serializing adapters implementing `redis.Serializer` to marshal and unmarshal cached Go objects.

---

## 1. Features

- **JSONSerializer**: Standard JSON format encoder.
- **GOBSerializer**: Optimized Go binary encoder.

---

## 2. API Examples

```go
import "prahari/shared/redis/serializer"

s := serializer.NewJSONSerializer()
bytesPayload, _ := s.Marshal(someStruct)
```
