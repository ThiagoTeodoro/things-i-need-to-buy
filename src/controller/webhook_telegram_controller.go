package controller

import (
	"fmt"
	"io"
	"net/http"

	"things.i.need.to.buy/src/service"
)

type MemoryMessages []service.Update

//Global declaration
var (
	instance MemoryMessages
)

//GetMainEntrypoint Function to responde the http get on '/telegram/webhook'
func TelegramWebhook(responseWriter http.ResponseWriter, request *http.Request) {

	//Insure the creation
	if instance == nil {
		instance = make(MemoryMessages, 0, 1000)
	}

	message, _ := service.ParseTelegramMessage(request)

	instance = append(instance, *message)

	io.WriteString(responseWriter, fmt.Sprint(instance))
}
