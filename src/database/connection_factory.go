package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

//Connect - To connect on database, don't forget to close the connection.
func Connect() sql.DB {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected on database!")
	return *db
}
