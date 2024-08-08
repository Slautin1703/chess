package main

import (
	"database/sql"
	"github.com/Slautin1703/games/shared/config"
	"log"
	"net/http"
	db2 "user-service/internal/infrastructure/db"
	httpInterfaces "user-service/internal/interfaces/http"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Initialize database connection
	pgDB, err := db2.ConnectToPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer func(pgDB *sql.DB) {
		err := pgDB.Close()
		if err != nil {
			log.Fatalf("could not close database connection: %v", err)
		}
	}(pgDB)

	db2.RunPostgresMigrations(cfg.DatabaseURL)

	// Set up routes
	router := httpInterfaces.NewRouter()

	corsRouter := httpInterfaces.CORSMiddleware(router)

	// Start the server
	serverAddr := cfg.ServerAddress
	log.Printf("Starting server on %s", serverAddr)
	if err := http.ListenAndServe(serverAddr, corsRouter); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
