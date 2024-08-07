package db

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func ConnectToMongoDB(uri string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")
	return &MongoDB{Client: client}, nil
}

func (db *MongoDB) Disconnect() {
	if err := db.Client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}
}
