package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func hello(responseWriter http.ResponseWriter, request *http.Request) {
	io.WriteString(responseWriter, "Hello World!")
}

func main() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", hello)

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, myRouter)
}
