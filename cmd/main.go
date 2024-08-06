package main

import (
	"awesomeProject/db"
	"awesomeProject/internal/handler"
	"database/sql"
	"log"
	"net/http"

	"awesomeProject/internal/config"
	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Initialize MongoDB
	mongoDB, err := db.ConnectToMongoDB(cfg.MongoURI)
	if err != nil {
		log.Fatalf("could not connect to MongoDB: %v", err)
	}
	defer mongoDB.Disconnect()

	// Initialize database connection
	pgDB, err := db.ConnectToPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer func(pgDB *sql.DB) {
		err := pgDB.Close()
		if err != nil {
			log.Fatalf("could not close database connection: %v", err)
		}
	}(pgDB)

	db.RunPostgresMigrations(cfg.DatabaseURL)

	// Set up repository
	//userRepo := repository.NewUserRepository(db)
	//gameRepo := repository.NewGameRepository(db)
	//
	//// Set up services
	//userService := service.NewUserService(userRepo)
	//gameService := service.NewGameService(gameRepo)
	//
	//// Set up handlers
	//userHandler := handler.NewUserHandler(userService)
	//gameHandler := handler.NewGameHandler(gameService)

	// Create a new router
	r := mux.NewRouter()

	r.HandleFunc("/", handler.HelloHandler).Methods("GET")

	// Define routes
	//r.HandleFunc("/signup", userHandler.SignUp).Methods("POST")
	//r.HandleFunc("/login", userHandler.Login).Methods("POST")
	//r.HandleFunc("/user", middleware.RequireAuth(userHandler.GetUser)).Methods("GET")
	//r.HandleFunc("/gamesHistory", middleware.RequireAuth(gameHandler.GetGames)).Methods("GET")
	//r.HandleFunc("/createGame", middleware.RequireAuth(gameHandler.CreateGame)).Methods("POST")

	// Start the server
	serverAddr := cfg.ServerAddress
	log.Printf("Starting server on %s", serverAddr)
	if err := http.ListenAndServe(serverAddr, r); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
