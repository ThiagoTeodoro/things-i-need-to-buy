package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"things.i.need.to.buy/src/database"
	"things.i.need.to.buy/src/router"
)

// App export
type App struct {
	Router mux.Router
}

/*
	Main code init of application
*/
func main() {

	app := App{
		Router: router.RouterConfig(),
	}

	//This port is requirement buy Heroku configs.
	port := os.Getenv("PORT")
	//port = "8080"

	database.Connect()

	//Create a goroutine
	server := make(chan bool)
	log.Printf("Server started at port %v", port)
	go log.Fatal(http.ListenAndServe(":"+port, &app.Router))
	<-server
}
