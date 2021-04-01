package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

//Connect - To connect on database, don't forget to close the connection.
func Connect() sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("db_host"),
		os.Getenv("db_port"),
		os.Getenv("db_user"),
		os.Getenv("db_pass"),
		os.Getenv("db_name"))

	db, err := sql.Open("postgres", psqlInfo)

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
