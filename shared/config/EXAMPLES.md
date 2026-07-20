# Examples: Configurations Loader

The following sections display standard implementation configurations.

---

## 1. Go Service Boot Code

```go
package main

import (
	"context"
	"log"

	"prahari/shared/config"
)

func main() {
	ctx := context.Background()

	// Parse configurations. Inject nil clients for SSM/Secrets if not enabled.
	cfg, err := config.Load(ctx, nil, nil)
	if err != nil {
		log.Fatalf("Fatal: Configuration build failed: %v", err)
	}

	log.Printf("Booted service: %s on environment: %s", cfg.ServiceName, cfg.Environment)
}
```

---

## 2. Local `.env` Configuration File

```ini
ENVIRONMENT=development
SERVICE_NAME=auth-service
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=secret
DB_NAME=prahari_db
SECURITY_ALLOWED_ORIGINS=http://localhost:3000
KAFKA_BROKERS=localhost:9092
KAFKA_CLIENT_ID=auth-service-dev
```
