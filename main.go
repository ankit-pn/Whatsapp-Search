package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func load_env() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func connectToMongo() (*mongo.Client, error) {
	mongoURI := fmt.Sprintf("mongodb://%s", os.Getenv("MONGO"))
	fmt.Printf("%s\n", mongoURI)
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		// log.Fatalf("Failed to connect to MongoDB: %v", err)
		return nil, err
	}
	// defer client.Disconnect(context.TODO())

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		// log.Fatalf("Failed to ping MongoDB: %v", err)
		return nil, err
	}

	// db := client.Database("whatsappLogs")

	return client, nil
}

func main() {
	load_env()
	client, err := connectToMongo()
	if err != nil {
		log.Fatalf("Failed to connect MongoDB: %v", err)
		return
	}
	defer client.Disconnect(context.TODO())

	// allsurveyors(client,"whatsappLogs")
	// getChatIds(client,"whatsappLogs","wsurveyor3")

	// es_test()
	// getChatIds(client,"whatsappLogs","wsurveyor3")

	// getCollections(client,"whatsappLogs")
	// getSamples(client,"whatsappLogs")
	// fetchChatListByParticipantID(client,"whatsappLogs","639c864c3893ae2d73795ab1")
	// fetchAllMessages(client, "whatsappLogs")
	// allsurveyors(client,"whatsappLogs")
	//  getchatinfo(client,"whatsappLogs","2TGq7fQO0qgqyYBopuejnGWjchS4zCX3Dx43Kiqk6uU=@c.us")
	// getdocumentcount(client,"whatsappLogs")
	// getdatabase(client)
	// getdata(client,"whatsappLogs","surveyors")
	// getdata(client,"whatsappLogs","messages")
	// getmsg(client,"whatsappLogs")
	// upgroup := [7]string{"wsurveyor1", "wsurveyor2", "wsurveyor3", "wsurveyor4", "wsurveyor5", "wsurveyor6", "ved"}
	// for _, value := range upgroup {
	// 	fetchChatListForOneSurveyor(client, "whatsappLogs", value)
	// }

	perform_index_up_group(client)

	//  getdata(client,"whatsappLogs","participants")
}
