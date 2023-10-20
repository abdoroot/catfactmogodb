package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type catFactServer struct {
	client *mongo.Client
}

func newServer(client *mongo.Client) *catFactServer {
	return &catFactServer{
		client: client,
	}
}

func (s catFactServer) handleGetAllFacts(w http.ResponseWriter, r *http.Request) {
	coll := s.client.Database("catfact").Collection("facts")
	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	results := []bson.M{} //map[string]interface{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

type catFactWorker struct {
	client *mongo.Client
}

func newWorker(client *mongo.Client) *catFactWorker {
	return &catFactWorker{
		client: client,
	}
}

func (w *catFactWorker) start() error {
	ticker := time.NewTicker(2 * time.Second)
	// collection
	collection := w.client.Database("catfact").Collection("facts")
	for {
		res, err := http.Get("https://catfact.ninja/fact")
		if err != nil {
			return err
		}
		catFact := make(map[string]interface{})
		decodeErr := json.NewDecoder(res.Body).Decode(&catFact)
		if decodeErr != nil {
			return decodeErr
		}
		collection.InsertOne(context.TODO(), catFact)
		<-ticker.C //block until we get this``
	}
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	worker := newWorker(client)
	go worker.start()

	server := newServer(client)
	http.HandleFunc("/facts", server.handleGetAllFacts)
	http.ListenAndServe(":3000", nil)
}
