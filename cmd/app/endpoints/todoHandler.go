package endpoints

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/VijoPlays/golang-demo/cmd/app/database"
)

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo database.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		fmt.Println(err)

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	//Check for http.Method
	switch r.Method {
	case http.MethodPost:

		database.AddTodo(todo)

		return
	case http.MethodPatch:

		err = database.UpdateTodo(todo)
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
