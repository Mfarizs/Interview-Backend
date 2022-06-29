package request

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Body struct {
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

func JsonDecode(r *http.Request) (Body, error) {
	var request Body
	errDecode := json.NewDecoder(r.Body).Decode(&request)

	return request, errDecode
}

func BodyDecode(r *http.Request) (Body, error) {
	var request Body
	errDecode := schema.NewDecoder().Decode(&request, r.PostForm)
	return request, errDecode
}

func QueryDecode(r *http.Request) (Body, error) {
	var request Body
	errDecode := schema.NewDecoder().Decode(&request, r.URL.Query())
	return request, errDecode
}
