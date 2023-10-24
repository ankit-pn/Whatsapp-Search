package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func getSamples(client *mongo.Client, dbname string) {

	db := client.Database(dbname)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// List all collections in the database
	collectionNames, err := db.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Failed to list collection names: %v", err)
	}

	// Iterate through each collection to fetch one document as a sample
	for _, collectionName := range collectionNames {
		if collectionName == "largeFiles.chunks" {
			continue // Skip the largeFiles.chunks collection
		}
		collection := db.Collection(collectionName)
		var sampleDoc bson.M
		err := collection.FindOne(ctx, bson.M{}).Decode(&sampleDoc)
		if err != nil {
			fmt.Printf("Failed to get a sample document from collection %s: %v\n", collectionName, err)
			continue
		}

		prettyJSON, err := json.MarshalIndent(sampleDoc, "", " ")
		if err != nil {
			fmt.Printf("Failed to marshal sample document to JSON %v\n", err)
		}
		fmt.Printf("Sample Document from collection %s : \n%s\n", collectionName, prettyJSON)
	}
}
