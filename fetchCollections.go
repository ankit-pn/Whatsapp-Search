package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func getCollections(client *mongo.Client, dbname string) {

	db := client.Database(dbname)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// List all collections in the whatsappLogs database
	collectionNames, err := db.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Failed to list collection names: %v", err)
	}

	fmt.Println("Collections in whatsappLogs database:", collectionNames)
}
