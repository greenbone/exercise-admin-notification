package rest

import (
	"admin-alarm/internal/adapter"
	"admin-alarm/internal/adapter/representation"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	log.Print("Starting Admin Notification service on http://localhost:8080/api/notify")

	alarmRouter := mux.NewRouter().StrictSlash(true)
	alarmRouter.HandleFunc("/api/notify", processAlarm).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", alarmRouter))
}

func processAlarm(writer http.ResponseWriter, request *http.Request) {
	headerContentType := request.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		errorMessage := representation.Error{}
		errorMessage.Message = "Content Type must be application/json"
		adapter.SendResponse(writer, http.StatusBadRequest, errorMessage)
		return
	}

	alarm := representation.Alarm{}
	err := parseData(request, &alarm)
	if err != nil {
		errorMessage := representation.Error{}
		errorMessage.Message = "Could not parse data"
		adapter.SendResponse(writer, http.StatusBadRequest, errorMessage)
		return
	}
	err = alarm.OK()
	if err != nil {
		adapter.SendResponse(writer, http.StatusUnprocessableEntity, err.Error())
		return
	}
	log.Printf("Notification successfully received: %+v\n", alarm)
	adapter.SendResponse(writer, http.StatusOK, alarm)
}

func parseData(request *http.Request, model interface{}) error {
	return json.NewDecoder(request.Body).Decode(model)
}
