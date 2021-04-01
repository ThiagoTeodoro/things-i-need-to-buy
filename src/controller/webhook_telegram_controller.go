package controller

import (
	"net/http"

	"things.i.need.to.buy/src/service"
)

//TelegramWebhook Function to responde the http get on '/telegram/webhook'
func TelegramWebhook(responseWriter http.ResponseWriter, request *http.Request) {

	message, _ := service.TelegramMessageParse(request)

	service.SendToTelegramChat(message.Message.Chat.ID, service.TelegramDecisionMaker(message.Message.Text))
}
