package controller

import (
	"io"
	"net/http"

	"things.i.need.to.buy/src/service"
)

//GetMainEntrypoint Function to responde the http get on '/'
func GetMainEntrypoint(responseWriter http.ResponseWriter, request *http.Request) {
	io.WriteString(responseWriter, service.HelloWord())
}
