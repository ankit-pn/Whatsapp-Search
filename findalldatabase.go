package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func getdatabase(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	databases, err := client.ListDatabaseNames(ctx, map[string]interface{}{})
	if err != nil {
		log.Fatalf("Failed to list database names: %v", err)
		return
	}

	fmt.Println("Available databases:")
	for _, dbName := range databases {
		fmt.Println(" - " + dbName)
	}
}
