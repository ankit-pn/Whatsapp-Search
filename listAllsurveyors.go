package main

import (
	// "fmt"
	// "fmt"
	"log"
	// // "bytes"
	"time"
	// // "encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	// // "go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	// // "io"
	"context"
)

type ContactInfoSurveyor struct {
	ID          string `bson:"_id"`
	Address     string `bson:"address"`
	Email       string `bson:"email"`
	PhoneNumber string `bson:"phoneNumber"`
}

type Surveyor struct {
	Version            int                 `bson:"__v"`
	ID                 string              `bson:"_id"`
	AddedBy            string              `bson:"addedBy"`
	Bio                string              `bson:"bio"`
	ContactInfo        ContactInfoSurveyor `bson:"contactInfo"`
	DateOfRegistration time.Time           `bson:"dateOfRegistration"`
	LastActiveAt       time.Time           `bson:"lastActiveAt"`
	Name               string              `bson:"name"`
	ParticipantsAdded  []string            `bson:"participantsAdded"`
	Password           string              `bson:"password"`
	RefreshToken       string              `bson:"refreshToken"`
	SurveyDisabled     bool                `bson:"surveyDisabled"`
	Username           string              `bson:"username"`
}

func allsurveyors(client *mongo.Client, dbname string) {

	db := client.Database(dbname)
	surveyorsCollection := db.Collection("surveyors")

	findOptions := options.Find()

	var results []*Surveyor

	curr, err := surveyorsCollection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for curr.Next(context.TODO()) {
		var elem Surveyor
		err := curr.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	curr.Close(context.TODO())

	for _, result := range results {
		// fmt.Printf("Surveyor %d:\n", i+1)
		// fmt.Printf("  ID: %s\n", result.ID)
		// fmt.Printf("  Name: %s\n", result.Name)
		// fmt.Printf("  Bio: %s\n", result.Bio)
		participants := result.ParticipantsAdded
		for _, value := range participants {
			fetchFileNameListByParticipantID(client, "whatsappLogs", value)
		}
		// fmt.Printf("  Username: %s\n", result.Username)
		// Add more fields as needed
		// fmt.Println("-----------------------------")
	}
}

func fetchFileNameListForOneSurveyor(client *mongo.Client, dbname string, surveyorUserName string) []string {
	db := client.Database(dbname)
	surveyorsCollection := db.Collection("surveyors")

	var result Surveyor
	filter := bson.M{"username": surveyorUserName}

	err := surveyorsCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}
	participantList := result.ParticipantsAdded

	var results []string
	for _, value := range participantList {
		fetchedList := fetchFileNameListByParticipantID(client, "whatsappLogs", value)
		results = append(results, fetchedList...)
	}

	return results
}
