package service

import "things.i.need.to.buy/src/dto"

//TelegramDecisionMaker - To response, and make actino by text send by user.
func TelegramDecisionMaker(msg dto.Message) string {

	switch msg.Text {
	case "linux":

		return "excelente os"
	default:

		return "Bem vindo " + msg.From.FirstName + " ao modulo Javis para armazenamento de lista, ainda estou em construção, em breve estarei em funcionamento."
	}
}
