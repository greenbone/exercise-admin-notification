package rest

import (
	"admin-alarm/internal/adapter/representation"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	alarmRouter := mux.NewRouter().StrictSlash(true)
	alarmRouter.HandleFunc("/notification", processAlarm).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", alarmRouter))
}

func processAlarm(writer http.ResponseWriter, request *http.Request) {
	alarm := representation.Alarm{}
	err := parseData(request, &alarm)
	if err != nil {
		errorMessage := representation.Error{}
		errorMessage.Message = "Could not parse data"
		sendResponse(writer, http.StatusBadRequest, errorMessage)
		return
	}
	err = alarm.OK()
	if err != nil {
		sendResponse(writer, http.StatusUnprocessableEntity, err.Error())
		return
	}
	log.Printf("Notification successful receiced: %+v\n", alarm)
	sendResponse(writer, http.StatusOK, alarm)
}

func sendResponse(writer http.ResponseWriter, statusCode int, body interface{}) {
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

func parseData(request *http.Request, model interface{}) error {
	return json.NewDecoder(request.Body).Decode(model)
}
