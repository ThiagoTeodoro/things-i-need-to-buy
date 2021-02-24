package dto

//Update Struct to recive the messages from Telegram this is the Main Struct
type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

//Message Struct to recive the messages from Telegram this is a secondary Struct
type Message struct {
	Text    string  `json:"text"`
	Chat    Chat    `json:"chat"`
	Contact Contact `json:"contact"`
}

//Chat Struct to recive the messages from Telegram this is a secondary Struct
type Chat struct {
	ID int `json:"id"`
}

//Contact Struct to revice the contact data from Telegram
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	UserID      int    `json:"user_id"`
	Vcard       string `json:"vcard"`
}
