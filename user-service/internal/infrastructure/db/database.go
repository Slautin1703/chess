package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var DB *sql.DB

func ConnectToPostgres(dsn string) (*sql.DB, error) {
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := DB.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to database")
	return DB, nil
}

type MongoDB struct {
	Client *mongo.Client
}
