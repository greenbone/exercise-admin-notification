package main

import (
	alarmController "admin-alarm/internal/adapter/rest"
)

func main() {
	alarmController.HandleRequest()
}
