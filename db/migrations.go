package db

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

// RunPostgresMigrations applies migrations to the PostgreSQL database.
func RunPostgresMigrations(databaseURL string) {

	// Create a new migrate instance
	m, err := migrate.New(
		"file://db/migrations",
		databaseURL,
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	// Apply migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("PostgreSQL migrations applied successfully")
}
