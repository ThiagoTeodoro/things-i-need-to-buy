package repository

import (
	"database/sql"
	"log"

	"things.i.need.to.buy/src/dao"
	"things.i.need.to.buy/src/database"
)

//InsertItem - To Insert itens on table listItens
func InsertItem(userName string, description string, value float64) int {

	var db sql.DB = database.Connect()

	sqlStatement := `
	  INSERT INTO tb_item (userName, description, value)
	  VALUES ($1, $2, $3)
	  RETURNING id`

	id := 0
	err := db.QueryRow(sqlStatement, userName, description, value).Scan(&id)
	if err != nil {
		log.Println(err)
	}

	return id
}

func GetAllItems(userName string) []dao.Item {

	var db = database.Connect()
	var result []dao.Item

	rows, err := db.Query("SELECT * FROM tb_item WHERE userName=$1", userName)
	if err != nil {
		// handle this error better than this
		log.Println(err)
	}
	defer rows.Close()
	defer db.Close()

	for rows.Next() {

		var item dao.Item

		err = rows.Scan(&item.ID, &item.UserName, &item.Description, &item.Value, &item.CreateDate)

		if err != nil {
			// handle this error
			log.Println(err)
		}

		result = append(result, item)
	}

	// get any error encountered during iteration
	err = rows.Err()

	if err != nil {
		log.Println(err)
	}

	return result
}
