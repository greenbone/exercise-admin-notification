package adapter

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendResponse(writer http.ResponseWriter, statusCode int, body interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	if body == nil {
		return
	}
	err := json.NewEncoder(writer).Encode(body)
	if err != nil {
		log.Println("Could not parse body", err)
	}
}
