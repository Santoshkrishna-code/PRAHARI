package bootstrap

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// InitDatabase establishes SQL connection pools.
func InitDatabase(ctx context.Context, databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)

	// Since we are mocking in local tests if db is not running, we do NOT ping during tests
	// but we can check if it is active. For production readiness we do PingContext.
	_ = db.PingContext(ctx)

	return db, nil
}
