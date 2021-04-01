package service

import (
	"strings"

	"things.i.need.to.buy/src/database"
	"things.i.need.to.buy/src/dto"
)

//TelegramDecisionMaker - To response, and make actino by text send by user.
func TelegramDecisionMaker(msg dto.Message) string {

	switch {

	case strings.EqualFold(msg.Text, "olá"):

		return "Bem vindo " + msg.From.FirstName + " ao modulo Javis para armazenamento de lista, ainda estou em construção, em breve estarei em funcionamento."

	case strings.HasPrefix(msg.Text, "/add "):

		database.Connect()

		//Sanatize
		// description := strings.ReplaceAll(msg.Text, "/add ", "")
		// parts := strings.Split(description, " ")
		// description = parts[0]
		// value, err := strconv.ParseFloat(parts[1], 64)
		// if err != nil {
		// 	panic(err)
		// }

		// return fmt.Sprintf("O item %d foi cadastro na nossa base de dados.", repository.InsertItem(msg.From.UserName, description, value))

	case strings.HasPrefix(msg.Text, "/delete "):

		return ""

	default:

		return "Que pena, não consegui te entender, vou registar pra melhorar minhas funções."

	}
}
