package response

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Body struct {
	ID          int       `json:"id" bson:"ids"`
	Category    int       `json:"category" bson:"category"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Footer      string    `json:"footer" bson:"footer"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
}

type erroMsg struct {
	Code    string `json:"code"`
	Message string `json:"errorMessage"`
}

func Send(w http.ResponseWriter, data Body) bool {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
		return true
	}

	log.Println(string(b))
	return true
}
