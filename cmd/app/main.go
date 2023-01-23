package main

import (
	"log"
	"net/http"
	"os"

	"github.com/VijoPlays/golang-demo/cmd/app/endpoints"
)

//When writing Go-Code, do not use empty packages. F.ex.: A package that only contains other packages (and no files itself) is pointless in Go and should be avoided.

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
