package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	// "sync"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"go.mongodb.org/mongo-driver/mongo"
)

// SearchResult is used to unmarshal Elasticsearch response
type SearchResult struct {
	Hits struct {
		Total struct {
			Value int
		}
		Hits []struct {
			Source map[string]interface{} `json:"_source"`
		}
	} `json:"hits"`
}

func esi() *elasticsearch.Client {
	elasticSearchAddress := os.Getenv("ES")
	cfg := elasticsearch.Config{
		Addresses: []string{
			elasticSearchAddress, // Replace with your Elasticsearch instance address
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	return es
}

func indexdoc(es *elasticsearch.Client) {
	// Indexing a document
	req := esapi.IndexRequest{
		Index:      "test_index",
		DocumentID: "1",
		Body:       strings.NewReader(`{"title":"Test", "content":"Elasticsearch and Go"}`),
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
}

func searchdocs(es *elasticsearch.Client) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "Test",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("test_index"),
		es.Search.WithBody(&buf),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	defer res.Body.Close()

	// Decode the JSON response
	var result SearchResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	// Print the ID and document source for each hit.
	for _, hit := range result.Hits.Hits {
		log.Printf("Source: %v", hit.Source)
	}
}

// deleteIndex deletes an Elasticsearch index
func deleteIndex(es *elasticsearch.Client, indexName string) error {
	// Create the request
	req := esapi.IndicesDeleteRequest{
		Index: []string{indexName},
	}

	// Perform the request
	res, err := req.Do(context.Background(), es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Check for errors in response
	if res.IsError() {
		return fmt.Errorf("Error: %s", res.String())
	}

	return nil
}


func indexMessage(es *elasticsearch.Client, indexName string,message interface{}){

	// defer wg.Done()
    
    messageJSON, err := json.Marshal(message)
    if err != nil {
        log.Printf("Error marshalling the message: %s", err)
        return
    }

    req := esapi.IndexRequest{
        Index:      indexName,
        Body:       strings.NewReader(string(messageJSON)),
        Refresh:    "true",
    }

    res, err := req.Do(context.Background(), es)
    if err != nil {
        log.Printf("Error indexing document: %s", err)
        return
    }
    defer res.Body.Close()
}


func perform_index_up_group(client *mongo.Client) {
    es := esi()
    // var wg sync.WaitGroup  // Declare a wait group

    upgroup := [7]string{"wsurveyor1", "wsurveyor2", "wsurveyor3", "wsurveyor4", "wsurveyor5", "wsurveyor6", "ved"}

    for _, value := range upgroup {
        fileList := fetchFileNameListForOneSurveyor(client, "whatsappLogs", value)
        
        for _, val := range fileList {
            messageList := fetchMessageByFileName(client, "whatsappLogs", val)

            for _, val := range messageList {
                // wg.Add(1)  // Increment the wait group counter
                indexMessage(es, "up_data", val)  // Launch a Goroutine
            }
        }
    }

    //  wg.Wait()  // Wait for all Goroutines to complete
}
