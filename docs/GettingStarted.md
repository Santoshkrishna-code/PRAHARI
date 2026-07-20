# Getting Started: Bootstrapping a Microservice

This guide walks you through bootstrapping a new microservice using the PRAHARI Shared Platform SDK.

---

## 1. Import Platform Packages

First, configure your service dependency to import the SDK:

```go
import (
	prahariConfig "prahari/shared/config"
	prahariLogger "prahari/shared/logger"
	prahariRedis "prahari/shared/redis"
	prahariMid "prahari/shared/middleware"
)
```

---

## 2. Load Configuration

Load settings dynamically from `.env` and environment variables:

```go
ctx := context.Background()
loader, _ := prahariConfig.NewLoader(prahariConfig.Options{})
var cfg MyConfig
_ = loader.Load(ctx, &cfg)
```

---

## 3. Setup Middlewares & Router

Compose request-handling chains:

```go
pipeline := prahariMid.New(
	requestid.Middleware,
	logging.Middleware,
	recovery.Middleware,
)
```
