package service

import (
	"encoding/json"
	"log"
	"net/http"

	"things.i.need.to.buy/src/dto"
)

//TelegramMessageParse function to parse the Telegram Messages
func TelegramMessageParse(request *http.Request) (*dto.Update, error) {
	var update dto.Update

	//To decode the Body from request on your struct Update
	if err := json.NewDecoder(request.Body).Decode(&update); err != nil {

		log.Printf("Error when the app try Parse the Telegram Message, %s", err.Error())
		return nil, err
	}

	return &update, nil
}
