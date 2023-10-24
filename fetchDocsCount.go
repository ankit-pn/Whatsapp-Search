package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func getdocumentcount(client *mongo.Client, dbname string) {
	db := client.Database(dbname)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collectionNames, err := db.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Failed to list Collections: %v", err)
	}
	fmt.Printf("Collections in %s database are:  %v \n", dbname, collectionNames)
	for _, collectionName := range collectionNames {
		collection := db.Collection(collectionName)
		count, err := collection.CountDocuments(ctx, bson.M{}, options.Count().SetMaxTime(2*time.Second))
		if err != nil {
			log.Printf("Failed to count document in collection %s: %v", collectionName, err)
			continue
		}

		fmt.Printf("Collection %s has %d documents. \n", collectionName, count)

	}
}
