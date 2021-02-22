package service

import (
	"encoding/json"
	"log"
	"net/http"
)

//Update Struct to recive the messages from Telegram this is the Main Struct
type Update struct {
	UpdateID int     `json: "update_id"`
	Message  Message `json: "message"`
}

//Message Struct to recive the messages from Telegram this is a secondary Struct
type Message struct {
	Text string `json: "text"`
	Chat Chat   `json: "chat"`
}

//Chat Struct to recive the messages from Telegram this is a secondary Struct
type Chat struct {
	Id int `json: "id"`
}

//ParseTelegramMessage function to parse the Telegram Messages
func ParseTelegramMessage(request *http.Request) (*Update, error) {
	var update Update

	//To decode the Body from request on your struct Update
	if err := json.NewDecoder(request.Body).Decode(&update); err != nil {

		log.Printf("Error when the app try Parse the Telegram Message, %s", err.Error())
		return nil, err
	}

	return &update, nil
}
