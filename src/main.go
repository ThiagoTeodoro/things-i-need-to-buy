package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
	port = "8080"
	http.ListenAndServe(":"+port, &app.Router)
}
