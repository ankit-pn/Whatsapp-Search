package main

import (
	"context"
	"fmt"
	// "fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ContactInfo struct {
	ID    string `bson:"_id"`
	Email string `bson:"email"`
}

type ConsentedChatUser struct {
	ID      string `bson:"0"`
	Consent bool   `bson:"1"`
}

type User struct {
	Version             int             `bson:"__v"`
	ID                  string          `bson:"_id"`
	AddedBy             string          `bson:"addedBy"`
	AddedByName         string          `bson:"addedByName"`
	AutoForward         bool            `bson:"autoforward"`
	AutoForwardReceiver string          `bson:"autoforward_receiver"`
	AutoForwardSender   string          `bson:"autoforward_sender"`
	Bio                 string          `bson:"bio"`
	ClientID            string          `bson:"clientId"`
	ClientStatus        string          `bson:"clientStatus"`
	ConsentedChatUsers  [][]interface{} `bson:"consentedChatUsers"` // or []ConsentedChatUser if possible
	ContactInfo         ContactInfo     `bson:"contactInfo"`
	DateOfRegistration  time.Time       `bson:"dateOfRegistration"` // or time.Time if you're going to parse it
	DaysLeft            int             `bson:"daysLeft"`
	IsLogging           bool            `bson:"isLogging"`
	IsRevoked           bool            `bson:"isRevoked"`
	Name                string          `bson:"name"`
	SurveyDisabled      bool            `bson:"surveyDisabled"`
}

func getChatIds(client *mongo.Client, dbname string, surveyorName string) {
	db := client.Database(dbname)

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	participantCollections := db.Collection("participants")

	filter := bson.M{"addedByName": surveyorName}

	findOptions := options.Find()

	var results []*User

	curr, err := participantCollections.Find(context.TODO(), filter, findOptions)

	if err != nil {
		log.Fatal(err)
	}

	for curr.Next(context.TODO()) {
		var elem User
		err := curr.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	curr.Close(context.TODO())

	var chatIDList []string
	for _, result := range results {
		participantChatIDList := result.ConsentedChatUsers
		// fmt.Printf("X: %v",participantChatIDList)
		for _, chatid := range participantChatIDList {
			if len(chatid) >= 2 {
				if chatID, ok := chatid[0].(string); ok {
					if consent, ok := chatid[1].(bool); ok && consent { // Check if chatid[1] is true
						chatIDList = append(chatIDList, chatID)
					}
				}
			}
		}
	}

	ans := 0
	for i, result := range chatIDList {
		fmt.Printf("id %v: %s\n", i, result)
		// ans+=getchatinfo(client,"whatsappLogs",string(result))
	}
	fmt.Print(ans)
	// for i, result := range results {
	//     fmt.Printf("User %d:\n", i+1)
	//     // fmt.Printf("  Version: %d\n", result.Version)
	//     // fmt.Printf("  ID: %s\n", result.ID)
	//     fmt.Printf("  AddedBy: %s\n", result.AddedBy)
	//     fmt.Printf("  AddedByName: %s\n", result.AddedByName)
	//     // fmt.Printf("  AutoForward: %v\n", result.AutoForward)
	//     // fmt.Printf("  Bio: %s\n", result.Bio)
	//     // fmt.Printf("  ClientID: %s\n", result.ClientID)
	//     // fmt.Printf("  ClientStatus: %s\n", result.ClientStatus)
	//     // fmt.Printf("  DateOfRegistration: %s\n", result.DateOfRegistration)
	//     // fmt.Printf("  DaysLeft: %d\n", result.DaysLeft)
	//     // fmt.Printf("  IsLogging: %v\n", result.IsLogging)
	//     // fmt.Printf("  IsRevoked: %v\n", result.IsRevoked)
	//     fmt.Printf("  Name: %s\n", result.Name)
	//     // fmt.Printf("  SurveyDisabled: %v\n", result.SurveyDisabled)
	//     // fmt.Printf("  Contact Info:\n")
	//     // fmt.Printf("    ID: %s\n", result.ContactInfo.ID)
	//     // fmt.Printf("    Email: %s\n", result.ContactInfo.Email)
	//     // Add more fields as needed
	//     fmt.Println("-----------------------------")
	// }
	// fmt.Printf("Found multiple documents: %+v\n", results)
}

func filterlist(client *mongo.Client, dbname string, surveyorName string) {
	db := client.Database(dbname)

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	participantCollections := db.Collection("participants")

	filter := bson.M{"addedByName": surveyorName}

	findOptions := options.Find()

	var results []*User

	curr, err := participantCollections.Find(context.TODO(), filter, findOptions)

	if err != nil {
		log.Fatal(err)
	}

	for curr.Next(context.TODO()) {
		var elem User
		err := curr.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	curr.Close(context.TODO())

	// var chatIDList []string
	// for _,result := range results{
	// 	participantChatIDList := result.ConsentedChatUsers
	// 	for _,chatid := range participantChatIDList{
	// 		chatIDList = append(chatIDList, chatid.ID)
	// 	}
	// }

	// for _,result:=range chatIDList{
	// 	fmt.Printf("id: %s", result)
	// }
	// for i, result := range results {
	//     fmt.Printf("User %d:\n", i+1)
	//     // fmt.Printf("  Version: %d\n", result.Version)
	//     // fmt.Printf("  ID: %s\n", result.ID)
	//     fmt.Printf("  AddedBy: %s\n", result.AddedBy)
	//     fmt.Printf("  AddedByName: %s\n", result.AddedByName)
	//     // fmt.Printf("  AutoForward: %v\n", result.AutoForward)
	//     // fmt.Printf("  Bio: %s\n", result.Bio)
	//     // fmt.Printf("  ClientID: %s\n", result.ClientID)
	//     // fmt.Printf("  ClientStatus: %s\n", result.ClientStatus)
	//     // fmt.Printf("  DateOfRegistration: %s\n", result.DateOfRegistration)
	//     // fmt.Printf("  DaysLeft: %d\n", result.DaysLeft)
	//     // fmt.Printf("  IsLogging: %v\n", result.IsLogging)
	//     // fmt.Printf("  IsRevoked: %v\n", result.IsRevoked)
	//     fmt.Printf("  Name: %s\n", result.Name)
	//     // fmt.Printf("  SurveyDisabled: %v\n", result.SurveyDisabled)
	//     // fmt.Printf("  Contact Info:\n")
	//     // fmt.Printf("    ID: %s\n", result.ContactInfo.ID)
	//     // fmt.Printf("    Email: %s\n", result.ContactInfo.Email)
	//     // Add more fields as needed
	//     fmt.Println("-----------------------------")
	// }
	// fmt.Printf("Found multiple documents: %+v\n", results)
}
