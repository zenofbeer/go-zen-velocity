package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// declare a new router
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	http.ListenAndServe(":8080", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	// pipe "Hello World" into the response writer
	fmt.Fprintf(writer, "Hello World!")
}
