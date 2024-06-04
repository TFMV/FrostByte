package polaris

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewClient creates a new database client for Polaris using pgx/v5.
func NewClient(host string, port int, user, password, dbname string) (*pgxpool.Pool, error) {
	// Data Source Name (DSN) format for PostgreSQL
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname)

	// Configure connection pool
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	// Connect to the database
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	return pool, nil
}
