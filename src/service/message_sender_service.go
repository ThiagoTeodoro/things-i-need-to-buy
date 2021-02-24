package service

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// SendToTelegramChat sends a text message to the Telegram chat identified by its chat Id
func SendToTelegramChat(chatID int, text string) {

	log.Printf("Sending message to chat_id: %d", chatID)
	var telegramAPI string = "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_BOT_TOKEN") + "/sendMessage"

	response, err := http.PostForm(
		telegramAPI,
		url.Values{
			"chat_id": {strconv.Itoa(chatID)},
			"text":    {text},
		},
	)

	//Check if http.PostForm produce erro.
	if err != nil {
		log.Printf("Error when the application try posting text to the chat: %s", err.Error())
	}

	//Reading the response
	defer response.Body.Close()
	var bodyBytes, errRead = ioutil.ReadAll(response.Body)

	if errRead != nil {
		log.Printf("Error in parsing telegram answer %s", errRead.Error())
	}

	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)
}
