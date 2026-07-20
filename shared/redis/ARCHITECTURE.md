# Architecture: Redis Platform SDK

This document details the design patterns, lock algorithms, and transaction rules governing the `shared/redis` package.

---

## 1. Key Design Tenets

- **Pool Management**: Uses connection pools with automated reconnections and health assertions.
- **Optimistic Concurrency**: Uses optimistic transactions via WATCH and Lua scripts to perform race-condition-free increments.
- **Watchdog Lease Extension**: Runs background go-routines to extend active lock keys, avoiding premature expires.

---

## 2. Lock Safety Mechanics

- **Safe lua unlock**: Checks value matches before deletion:
  ```lua
  if redis.call("get", KEYS[1]) == ARGV[1] then
      return redis.call("del", KEYS[1])
  else
      return 0
  end
  ```
- **Redlock Quorum**: Requires locking at least `n/2 + 1` independent nodes, rolling back acquisitions on quorum failure.
