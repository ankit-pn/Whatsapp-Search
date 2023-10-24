package main

import (
	"fmt"
	"log"
	// "bytes"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"time"
)

// type MediaData struct {
// 	Filename string `bson:"filename"`
// 	Mimetype string `bson:"mimetype"`
// 	Filesize int    `bson:"filesize"`
// }

// type Link struct {
// 	Link         string `bson:"link"`
// 	IsSuspicious bool   `bson:"isSuspicious"`
// }

// type Participant struct {
// 	Server     string `bson:"server"`
// 	User       string `bson:"user"`
// 	Serialized string `bson:"_serialized"`
// }

// type ID struct {
// 	FromMe      bool        `bson:"fromMe"`
// 	Remote      string      `bson:"remote"`
// 	ID          string      `bson:"id"`
// 	Participant Participant `bson:"participant"`
// 	Serialized  string      `bson:"_serialized"`
// }

// type FromToAuthor struct {
// 	Server     string `bson:"server"`
// 	User       string `bson:"user"`
// 	Serialized string `bson:"_serialized"`
// }

// type Data struct {
// 	ID                          ID           `bson:"id"`
// 	RowID                       int64        `bson:"rowId"`
// 	Body                        string       `bson:"body"`
// 	Type                        string       `bson:"type"`
// 	T                           int64        `bson:"t"`
// 	From                        FromToAuthor `bson:"from"`
// 	To                          FromToAuthor `bson:"to"`
// 	Author                      FromToAuthor `bson:"author"`
// 	Ack                         int          `bson:"ack"`
// 	Invis                       bool         `bson:"invis"`
// 	Star                        bool         `bson:"star"`
// 	KicNotified                 bool         `bson:"kicNotified"`
// 	InteractiveAnnotations      []string     `bson:"interactiveAnnotations"`
// 	DeprecatedMms3Url           string       `bson:"deprecatedMms3Url"`
// 	DirectPath                  string       `bson:"directPath"`
// 	Mimetype                    string       `bson:"mimetype"`
// 	Duration                    string       `bson:"duration"`
// 	Filehash                    string       `bson:"filehash"`
// 	Size                        int64        `bson:"size"`
// 	MediaKey                    string       `bson:"mediaKey"`
// 	MediaKeyTimestamp           int64        `bson:"mediaKeyTimestamp"`
// 	IsViewOnce                  bool         `bson:"isViewOnce"`
// 	Width                       int          `bson:"width"`
// 	Height                      int          `bson:"height"`
// 	StaticUrl                   string       `bson:"staticUrl"`
// 	IsFromTemplate              bool         `bson:"isFromTemplate"`
// 	PollOptions                 []string     `bson:"pollOptions"`
// 	PollInvalidated             bool         `bson:"pollInvalidated"`
// 	LatestEditMsgKey            string       `bson:"latestEditMsgKey"`
// 	LatestEditSenderTimestampMs string       `bson:"latestEditSenderTimestampMs"`
// 	Broadcast                   bool         `bson:"broadcast"`
// 	MentionedJidList            []string     `bson:"mentionedJidList"`
// 	IsVcardOverMmsDocument      bool         `bson:"isVcardOverMmsDocument"`
// 	IsForwarded                 bool         `bson:"isForwarded"`
// 	ForwardingScore             int          `bson:"forwardingScore"`
// 	Labels                      []string     `bson:"labels"`
// 	HasReaction                 bool         `bson:"hasReaction"`
// 	EphemeralOutOfSync          bool         `bson:"ephemeralOutOfSync"`
// 	ProductHeaderImageRejected  bool         `bson:"productHeaderImageRejected"`
// 	LastPlaybackProgress        int          `bson:"lastPlaybackProgress"`
// 	IsDynamicReplyButtonsMsg    bool         `bson:"isDynamicReplyButtonsMsg"`
// 	IsMdHistoryMsg              bool         `bson:"isMdHistoryMsg"`
// 	StickerSentTs               int64        `bson:"stickerSentTs"`
// 	IsAvatar                    bool         `bson:"isAvatar"`
// 	RequiresDirectConnection    bool         `bson:"requiresDirectConnection"`
// 	PttForwardedFeaturesEnabled bool         `bson:"pttForwardedFeaturesEnabled"`
// 	IsEphemeral                 bool         `bson:"isEphemeral"`
// 	IsStatusV3                  bool         `bson:"isStatusV3"`
// 	Links                       []string     `bson:"links"`
// }

// type Message struct {
// 	MediaKey        string    `bson:"mediaKey"`
// 	Ack             int       `bson:"ack"`
// 	HasMedia        bool      `bson:"hasMedia"`
// 	Body            string    `bson:"body"`
// 	Type            string    `bson:"type"`
// 	Timestamp       int       `bson:"timestamp"`
// 	From            string    `bson:"from"`
// 	To              string    `bson:"to"`
// 	Author          string    `bson:"author"`
// 	DeviceType      string    `bson:"deviceType"`
// 	IsForwarded     bool      `bson:"isForwarded"`
// 	ForwardingScore int       `bson:"forwardingScore"`
// 	IsStatus        bool      `bson:"isStatus"`
// 	IsStarred       bool      `bson:"isStarred"`
// 	Broadcast       bool      `bson:"broadcast"`
// 	FromMe          bool      `bson:"fromMe"`
// 	HasQuotedMsg    bool      `bson:"hasQuotedMsg"`
// 	VCards          []string  `bson:"vCards"`
// 	MentionedIds    interface{}  `bson:"mentionedIds"`
// 	IsGif           bool      `bson:"isGif"`
// 	IsEphemeral     bool      `bson:"isEphemeral"`
// 	Links           []Link    `bson:"links"`
// 	MediaData       MediaData `bson:"mediaData"`
// 	Duration        string    `bson:"duration,omitempty"`
// 	MediaDownloaded bool      `bson:"mediaDownloaded"`
// 	Data            Data      `bson:"_data,omitempty"`
// 	ID              ID        `bson:"id,omitempty"`
// 	MsgID           string    `bson:"msg_id"`
// 	IsAnonymized    bool      `bson:"isAnonymized"`
// 	Title           string    `bson:"title,omitempty"`
// 	Description     string    `bson:"description,omitempty"`
// }

// type Media struct {
// 	Filename        string `bson:"filename"`
// 	Mimetype        string `bson:"mimetype"`
// 	Filesize        int    `bson:"filesize"`
// 	MediaDownloaded bool   `bson:"mediaDownloaded,omitempty"`
// }

func getmsg(client *mongo.Client, dbname string) {
	// database := client.Database(dbname)
	bucketOptions := options.GridFSBucket().SetName("largeFiles")
	bucket, err := gridfs.NewBucket(
		client.Database("whatsappLogs"),
		bucketOptions,
	)

	if err != nil {
		log.Fatalf("Could not create bucket: %v", err)
	}
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(dbname)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	messagesCollection := db.Collection("messages")
	cursor, err := messagesCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Failed to find messages: %v", err)
	}
	defer cursor.Close(ctx)
	count := 0

	for cursor.Next(context.TODO()) {
		var messagesRow struct {
			Messages struct {
				Length   int    `bson:"length"`
				Filename string `bson:"filename"`
			} `bson:"messages"`
		}
		err := cursor.Decode(&messagesRow)
		if err != nil {

			log.Fatalf("Error is there: %v", err)
		}

		if messagesRow.Messages.Length > 0 {
			// Get the file from GridFS
			fmt.Println(messagesRow.Messages.Filename)
			fileStream, err := bucket.OpenDownloadStreamByName(messagesRow.Messages.Filename)

			if err != nil {
				log.Fatal(err)
			}
			defer fileStream.Close()

			content, err := io.ReadAll(fileStream)
			if err != nil {
				log.Println("Error reading file:", err)
				continue
			}
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// if bytesRead != len(content) {
			// 	log.Fatalf("Expected to read %d bytes but read %d", len(content), bytesRead)
			// }

			var msgData []interface{}
			// fmt.Println("Attempting to parse:", string(content))

			err = json.Unmarshal(content, &msgData)
			if len(content) == 0 {
				log.Println("Warning: Attempting to parse empty content")
				continue
			}

			if err != nil {
				log.Fatal(err)
				return
			}

			prettyJSON, err := json.MarshalIndent(msgData, "", "  ")
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Println(string(prettyJSON))

		}
		count += 1
		// if (count>=1){
		// break
		// }
	}
}

func fetchMessageByFileName(client *mongo.Client, dbname string, Filename string) []interface{} {
	bucketOptions := options.GridFSBucket().SetName("largeFiles")
	bucket, err := gridfs.NewBucket(
		client.Database("whatsappLogs"),
		bucketOptions,
	)

	if err != nil {
		log.Fatalf("Could not create bucket: %v", err)
	}
	if err != nil {
		log.Fatal(err)
	}

	fileStream, err := bucket.OpenDownloadStreamByName(Filename)

	if err != nil {
		log.Fatal(err)
	}
	defer fileStream.Close()

	content, err := io.ReadAll(fileStream)
	if err != nil {
		log.Println("Error reading file:", err)
	}

	var msgData []interface{}
	// fmt.Println("Attempting to parse:", string(content))

	err = json.Unmarshal(content, &msgData)
	if len(content) == 0 {
		log.Println("Warning: Attempting to parse empty content")
	}

	if err != nil {
		log.Fatal(err)
	}

	// prettyJSON, err := json.MarshalIndent(msgData, "", "  ")
	// fmt.Println(string(prettyJSON))

	return msgData
}
