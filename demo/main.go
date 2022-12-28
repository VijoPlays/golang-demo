package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"vijo.golang-demo/demo/database"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.HandleFunc("/todo", todoHandler)
	http.HandleFunc("/todoList", todoHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		fmt.Print("pewpew")

		var model database.DatabaseModel
		var todo database.Todo
		err := json.NewDecoder(r.Body).Decode(&model)
		if err != nil {
			fmt.Print("yeet")
		}
		err = json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			fmt.Print("yeet")
		}
		fmt.Print(model)
		fmt.Print(todo)

		fmt.Print(err)
		// err := database.AddTodo()
		//TODO: Get stuff from body and then call repo
	case http.MethodPatch:

	default:
		err := errors.New("Method not available")
		w.Header().Add("error", err.Error())
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Print(err)
	}
}

func todoListHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: Copy todoHandler
}

//TODO: Rest endpoints
//TODO: Tests
// - REST endpoints
// 	Add Todo
// 	Update Todo (mark as done)
// 	Add TodoList
// 	Update (add) Todo to TodoList
// - Todo Modell (Titel, Done)
// - TodoList Modell (Titel, Liste*-*Todo)
// - InMemory
// - Unit Tests
