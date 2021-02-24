package dto

//Update Struct to recive the messages from Telegram this is the Main Struct
type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
	From     From    `json:"from"`
}

//From struct to recive data about the user.
type From struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	UserName  string `json:"username"` //Unic on intire Telegram
}

//Message Struct to recive the messages from Telegram this is a secondary Struct
type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

//Chat Struct to recive the messages from Telegram this is a secondary Struct
type Chat struct {
	ID int `json:"id"`
}
