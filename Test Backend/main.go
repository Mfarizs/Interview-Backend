package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Person struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IDs         int                `json:"id,omitempty" bson:"ids,omitempty"`
	Category    int                `json:"category,omitempty" bson:"category,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Footer      string             `json:"footer,omitempty" bson:"footer,omitempty"`
	Tags        []string           `json:"tags,omitempty" bson:"tags,omitempty"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
}

type People struct {
	ID          int       `json:"id,omitempty" bson:"ids,omitempty"`
	Title       string    `json:"title,omitempty" bson:"title,omitempty"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	Footer      string    `json:"footer,omitempty" bson:"footer,omitempty"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
}

type Input struct {
	ID        int       `json:"id,omitempty" bson:"ids,omitempty"`
	Category  int       `json:"category,omitempty" bson:"category,omitempty"`
	Items     []Item    `json:"items,omitempty" bson:"items,omitempty"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

type Item struct {
	Title       string `json:"title,omitempty" bson:"title,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Footer      string `json:"footer,omitempty" bson:"footer,omitempty"`
}

type Output struct {
	ID          int       `json:"id,omitempty" bson:"ids,omitempty"`
	Category    int       `json:"category,omitempty" bson:"category,omitempty"`
	Title       string    `json:"title,omitempty" bson:"title,omitempty"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	Footer      string    `json:"footer,omitempty" bson:"footer,omitempty"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
}

// func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {}
// func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) { }
// func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) { }

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/Add", CreateDataEndpoint).Methods("POST")
	router.HandleFunc("/backend/question/one", GetTestNumber1).Methods("GET")
	router.HandleFunc("/backend/question/two", GetTestNumber2).Methods("GET")
	router.HandleFunc("/backend/question/three", PostTestNumber3).Methods("POST")
	http.ListenAndServe(":8000", router)
}

func CreateDataEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person Person
	person.CreatedAt = time.Now()
	_ = json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("domain").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}

func GetTestNumber1(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var people []People
	collection := client.Database("domain").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	opts := options.Find().SetProjection(bson.M{"_id": 0})
	cursor, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person People
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(people)

}

func GetTestNumber2(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var people Person
	collection := client.Database("domain").Collection("people")
	filter := bson.D{{"description", bson.D{{"$regex", "Ergonomic"}}}, {"title", bson.D{{"$regex", "Ergonomic"}}}, {"tags", "Sports"}}
	opts := options.Find().SetSort(bson.D{{"ids", -1}}).SetLimit(3)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	result := make([]Person, 0)
	for cursor.Next(ctx) {
		err := cursor.Decode(&people)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, people)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)

}

func PostTestNumber3(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var warehouse Input
	warehouse.CreatedAt = time.Now()
	_ = json.NewDecoder(request.Body).Decode(&warehouse)
	wData := warehouse.Items
	loop := []Output{}
	for _, s := range wData {
		loop = append(loop, Output{
			ID:          warehouse.ID,
			Category:    warehouse.Category,
			Title:       s.Title,
			Description: s.Description,
			Footer:      s.Footer,
			CreatedAt:   warehouse.CreatedAt,
		})
	}
	json.NewEncoder(response).Encode(loop)

}
