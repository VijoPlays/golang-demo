package main

import (
	"log"
	"net/http"
	"os"

	"github.com/VijoPlays/golang-demo/src/endpoints"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.HandleFunc("/todo", endpoints.TodoHandler)
	http.HandleFunc("/todoList", endpoints.TodoListHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
