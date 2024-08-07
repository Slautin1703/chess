package main

import (
	"database/sql"
	httpInterfaces "game-service/internal/delivery/http"
	db2 "game-service/internal/infrastructure/db"
	"game-service/internal/infrastructure/redis"
	"github.com/Slautin1703/games/shared/config"
	"log"
	"net/http"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Initialize MongoDB
	mongoDB, err := db2.ConnectToMongoDB(cfg.MongoURI)
	if err != nil {
		log.Fatalf("could not connect to MongoDB: %v", err)
	}
	defer mongoDB.Disconnect()

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

	rdb := redis.NewRedisClient()
	defer func() {
		if err := rdb.Close(); err != nil {
			log.Fatalf("Failed to close Redis client: %v", err)
		}
	}()

	//db2.RunPostgresMigrations(cfg.DatabaseURL)

	router := httpInterfaces.NewRouter()

	corsRouter := httpInterfaces.CORSMiddleware(router)

	// Start the server
	serverAddr := cfg.ServerAddress
	log.Printf("Starting server on %s", serverAddr)
	if err := http.ListenAndServe(serverAddr, corsRouter); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
