package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ChatID struct {
	Serialized string `bson:"_serialized"`
	Server     string `bson:"server"`
	User       string `bson:"user"`
}

type Messages struct {
	Filename string `bson:"filename"`
	Length   int    `bson:"length"`
}

type ChatData struct {
	V             int       `bson:"__v"`
	ID            string    `bson:"_id"`
	ChatID        ChatID    `bson:"chatID"`
	ChatName      string    `bson:"chatName"`
	IsAnonymized  *bool     `bson:"isAnonymized"` // Using a pointer to bool to allow null value
	IsGroup       bool      `bson:"isGroup"`
	LastUpdated   time.Time `bson:"lastUpdated"` // could also be time.Time depending on how you want to handle it
	Messages      Messages  `bson:"messages"`
	ParticipantID string    `bson:"participantID"`
	Timestamp     int64     `bson:"timestamp"`
	UserID        string    `bson:"userID"`
	UserName      string    `bson:"userName"`
}

func fetchAllMessages(client *mongo.Client, dbname string) {
	db := client.Database(dbname)
	messagesCollection := db.Collection("messages")

	findOptions := options.Find()

	cur, err := messagesCollection.Find(context.Background(), bson.D{{}}, findOptions)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer cur.Close(context.Background())

	var results []ChatData
	if err := cur.All(context.Background(), &results); err != nil {
		fmt.Println("Error reading cursor:", err)
		return
	}
	var result bson.M
	err1 := messagesCollection.FindOne(context.Background(), bson.D{{}}).Decode(&result)
	if err1 != nil {
		fmt.Println("Error fetching a single document:", err1)
		return
	}
	fmt.Printf("Raw result: %+v\n", result)

	for i, result := range results {

		fmt.Printf("Message %d:\n", i+1)
		fmt.Printf("  ID: %s\n", result.ID)
		fmt.Printf("  ChatName: %s\n", result.ChatName)
		fmt.Printf("  _serialized: %s\n", result.ChatID.Serialized)
		fmt.Printf("  Server: %s\n", result.ChatID.Server)
		fmt.Printf("  User: %s\n", result.ChatID.User)
		fmt.Printf("  LastUpdated: %s\n", result.LastUpdated)
		// Add more fields as needed
		fmt.Println("-----------------------------")
	}
}

func listAllMessages(client *mongo.Client, dbName string) {
	ctx := context.Background()

	// Get database and collection
	db := client.Database(dbName)
	messagesCollection := db.Collection("messages")

	// Create a cursor for the query
	cur, err := messagesCollection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Error setting up cursor:", err)
		return
	}
	defer cur.Close(ctx)

	// Decode each document one at a time
	for cur.Next(ctx) {
		var elem ChatData
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("Error decoding:", err)
			return
		}

		// Print the _serialized field
		fmt.Println("_serialized:", elem.ChatID.Serialized)
	}

	// Check for errors from iterating over rows.
	if err := cur.Err(); err != nil {
		fmt.Println("Cursor iteration error:", err)
	}
}
