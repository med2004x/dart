package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func openDatabase(ctx context.Context) (*sql.DB, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	database, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("open database handle: %w", err)
	}

	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(5)
	database.SetConnMaxLifetime(30 * time.Minute)

	pingContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if err := database.PingContext(pingContext); err != nil {
		database.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return database, nil
}

func main() {
	database, err := openDatabase(context.Background())
	if err != nil {
		fmt.Println("database startup failed:", err)
		return
	}
	defer database.Close()

	fmt.Println("database connection verified")
}
