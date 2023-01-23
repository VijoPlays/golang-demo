package endpoints

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/VijoPlays/golang-demo/cmd/app/database"
)

func TodoListHandler(w http.ResponseWriter, r *http.Request) {
	//When creating and updating lists right now, the todo components are not being saved -

	var list database.TodoList
	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		fmt.Println(err)

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	//Check for http.Method
	switch r.Method {
	case http.MethodPost:

		database.AddTodoList(list)

		return
	case http.MethodPatch:

		err = database.UpdateTodoList(list)
		if err != nil {
			fmt.Println(err)

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		return

	default:
		err := errors.New("Method not available. Use Post or Patch instead.")
		fmt.Print(err)

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(err.Error()))
		return
	}
}
