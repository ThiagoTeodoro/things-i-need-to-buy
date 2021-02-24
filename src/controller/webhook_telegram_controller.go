package controller

import (
	"fmt"
	"io"
	"net/http"

	"things.i.need.to.buy/src/service"
)

//TelegramWebhook Function to responde the http get on '/telegram/webhook'
func TelegramWebhook(responseWriter http.ResponseWriter, request *http.Request) {

	message, _ := service.TelegramMessageParse(request)
	text := "Bem vindo ao modulo Javis para armazenamento de lista, ainda estou em construção, em breve estarei em funcionamento."

	service.SendToTelegramChat(message.Message.Chat.ID, text)

	io.WriteString(responseWriter, fmt.Sprint(text))
}
