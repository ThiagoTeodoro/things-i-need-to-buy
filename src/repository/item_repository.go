package repository

import (
	"database/sql"

	"things.i.need.to.buy/src/database"
)

//InsertItem - To Insert itens on table listItens
func InsertItem(userName, description string, value float64) int {

	var db sql.DB = database.Connect()

	sqlStatement := `
	  INSERT INTO tb_item (userName, description, value)
	  VALUES ($1, $2, $3)
	  RETURNING id`

	id := 0
	err := db.QueryRow(sqlStatement, userName, description, value).Scan(&id)
	if err != nil {
		panic(err)
	}

	return id
}
