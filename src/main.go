package main

import (
	"io"
	"net/http"
	"os"
)

func hello(responseWriter http.ResponseWriter, request *http.Request) {
	io.WriteString(responseWriter, "Hello World!")
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+port, nil)
}
