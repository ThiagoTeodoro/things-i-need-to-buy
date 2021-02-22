package router

import (
	"github.com/gorilla/mux"
	"things.i.need.to.buy/src/controller"
)

// RouterConfig Function with the all the Routers of the application
func RouterConfig() mux.Router {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/telegram/webhook", controller.TelegramWebhook)
	myRouter.HandleFunc("/", controller.GetMainEntrypoint)

	return *myRouter
}
