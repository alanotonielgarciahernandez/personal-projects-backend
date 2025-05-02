// dbConnection.go
// Database ping and connection functions.

package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection object to MongoDB Database.
var MongoCN = DBConnect()

// Define MongoDB Cluster direction.
var clientOptions = options.Client().ApplyURI(os.Getenv("DB_URI"))

// Connects to the MongoDB Database using the client options.
// Returns the database client.
func DBConnect() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	return client
}

// Pings the MongoDB Database.
func CheckDBConnection() error {
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return err
	}
	return nil
}
