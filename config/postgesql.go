package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DBPool *pgxpool.Pool

func InitPostgresDatabase(cfg *Provider) *pgxpool.Pool {
	// Create a connection pool configuration
	config, err := pgxpool.ParseConfig(cfg.Postgresql.DBConnection)
	if err != nil {
		log.Fatalf("Unable to parse database config: %v", err)
	}

	// Set connection pool configurations
	config.MaxConns = int32(cfg.Postgresql.DBMaxConns)
	config.MaxConnLifetime = time.Duration(cfg.Postgresql.DBMaxConnLifetime) * time.Minute
	config.MaxConnIdleTime = time.Duration(cfg.Postgresql.DBMaxConnIdletime) * time.Minute

	// Create connection pool
	DBPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}

	// Test the connection
	if err := DBPool.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	return DBPool
}

func ClosedDB() {
	if DBPool != nil {
		fmt.Println("Database connection closed")
		DBPool.Close()
	}
}
