package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
 
func ConnectDB() *mongo.Client {
	Mongo_URL := "mongodb://localhost:27017/"
	client, err := mongo.NewClient(options.Client().ApplyURI(Mongo_URL))
 
	if err != nil {
		log.Fatal(err)
	}
 
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	err = client.Connect(ctx)
	defer cancel()
 
	if err != nil {
		log.Fatal(err)
	}
 
	fmt.Println("Connected to mongoDB")
	return client
}