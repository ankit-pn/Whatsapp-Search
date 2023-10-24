package main

import (
	"context"
	// "fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

// type ChatID struct {
// 	Serialized string `bson:"_serialized"`
// 	Server     string `bson:"server"`
// 	User       string `bson:"user"`
// }

// type Messages struct {
// 	Filename string `bson:"filename"`
// 	Length   int    `bson:"length"`
// }

//	type ChatData struct {
//		V             int       `bson:"__v"`
//		ID            string    `bson:"_id"`
//		ChatID        ChatID    `bson:"chatID"`
//		ChatName      string    `bson:"chatName"`
//		IsAnonymized  *bool     `bson:"isAnonymized"` // Using a pointer to bool to allow null value
//		IsGroup       bool      `bson:"isGroup"`
//		LastUpdated   time.Time `bson:"lastUpdated"` // could also be time.Time depending on how you want to handle it
//		Messages      Messages  `bson:"messages"`
//		ParticipantID string    `bson:"participantID"`
//		Timestamp     int64     `bson:"timestamp"`
//		UserID        string    `bson:"userID"`
//		UserName      string    `bson:"userName"`
//	}
func fetchFileNameListByParticipantID(client *mongo.Client, dbname string, participantId string) []string {
	db := client.Database(dbname)
	chatCollections := db.Collection("messages")

	objID, err := primitive.ObjectIDFromHex(participantId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"participantID": objID}

	findOptions := options.Find()

	var results []*ChatData
	curr, err := chatCollections.Find(context.TODO(), filter, findOptions)

	if err != nil {
		log.Fatal(err)
	}

	for curr.Next(context.TODO()) {
		var elem ChatData
		err := curr.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	var fileNameList []string

	for _, value := range results {
		fileNameList = append(fileNameList, value.Messages.Filename)
	}

	return fileNameList
}
