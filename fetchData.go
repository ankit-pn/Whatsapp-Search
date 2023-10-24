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

func getdata(client *mongo.Client, dbname string, collectionName string) {
	db := client.Database(dbname)

	collection := db.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Failed to find documents: %v", err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		prettyJSON, err := json.MarshalIndent(result, "", " ")
		if err != nil {
			log.Fatalf("Failed to marshal to JSON: %v", err)
		}

		fmt.Printf("Document: \n%s\n", prettyJSON)

	}
	// Check for errors from iterating the cursor
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
}
