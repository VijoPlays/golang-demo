package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VijoPlays/golang-demo/cmd/app/database"
	"github.com/VijoPlays/golang-demo/cmd/app/endpoints"
)

//TODO: Get endpoint for assertion
//FIXME: Usually you'd test the packages, not main.go - wanted to use the httptest server here

// setUp is a beforeEach function and needs to be run before every test to ensure correct behaviour
// t.Run() starts a new GoRoutine: Need to run setUp() in every routine itself, else you might encounter a data race
func setUp() {
	database.CleanUpDatabase()
}

func TestTodoEndpointsSuccess(t *testing.T) {
	t.Run("adding a new todo object succeeds", func(t *testing.T) {
		setUp()
		//Arrange
		todo := database.Todo{Title: "testTitle", Done: true, ID: "0"}
		var buffer bytes.Buffer
		err := json.NewEncoder(&buffer).Encode(todo)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}

		request, err := http.NewRequest(http.MethodPost, "/todo", &buffer)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}
		response := httptest.NewRecorder()

		//Act
		endpoints.TodoHandler(response, request)

		//Assert
		statusCode := response.Result().StatusCode
		if http.StatusOK != statusCode {
			t.Errorf("assertion failed. expected: '%v', actual: '%v'", http.StatusOK, statusCode)
		}
		if database.TodoSlice.Exists(todo) < 0 {
			t.Errorf("assertion failed. object was not created properly in database")
		}
	})

	t.Run("updating an existing todo object succeeds", func(t *testing.T) {
		setUp()
		//Arrange
		todo := database.Todo{Title: "testTitle", Done: true, ID: "0"}
		var buffer bytes.Buffer
		err := json.NewEncoder(&buffer).Encode(todo)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}
		//Add object to database
		request, err := http.NewRequest(http.MethodPost, "/todo", &buffer)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}
		response := httptest.NewRecorder()
		endpoints.TodoHandler(response, request)
		//Prepare patch request
		err = json.NewEncoder(&buffer).Encode(todo)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}
		request, err = http.NewRequest(http.MethodPatch, "/todo", &buffer)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}
		response = httptest.NewRecorder()

		//Act
		endpoints.TodoHandler(response, request)

		//Assert
		statusCode := response.Result().StatusCode
		if http.StatusOK != statusCode {
			t.Errorf("assertion failed. expected: '%v', actual: '%v'", http.StatusOK, statusCode)
		}
		if database.TodoSlice.Exists(todo) != 0 {
			t.Errorf("assertion failed. object was not created properly in database")
		}
	})
}

func TestTodoEndpointsFailure(t *testing.T) {
	t.Run("adding an existing todo object results in a BadRequest error", func(t *testing.T) {
		setUp()
		//Arrange
		todo := database.Todo{Title: "testTitle", Done: true, ID: "0"}
		var buffer bytes.Buffer
		err := json.NewEncoder(&buffer).Encode(todo)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}

		request, err := http.NewRequest(http.MethodPost, "/todo", &buffer)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}
		response := httptest.NewRecorder()

		//Act
		endpoints.TodoHandler(response, request)
		//Resend the same request
		request, err = http.NewRequest(http.MethodPost, "/todo", &buffer)
		if err != nil {
			t.Errorf("failed to execute act block of the test: %v", err)
		}
		endpoints.TodoHandler(response, request)

		//Assert
		statusCode := response.Result().StatusCode
		if http.StatusBadRequest != statusCode {
			t.Errorf("assertion failed. expected: '%v', actual: '%v'", http.StatusBadRequest, statusCode)
		}
	})

	t.Run("updating a non existing todo object results in a BadRequest error", func(t *testing.T) {
		setUp()
		//Arrange
		todo := database.Todo{Title: "testTitle", Done: true, ID: "0"}
		var buffer bytes.Buffer
		err := json.NewEncoder(&buffer).Encode(todo)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}

		request, err := http.NewRequest(http.MethodPatch, "/todo", &buffer)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}
		response := httptest.NewRecorder()

		//Act
		endpoints.TodoHandler(response, request)

		//Assert
		statusCode := response.Result().StatusCode
		if http.StatusBadRequest != statusCode {
			t.Errorf("assertion failed. expected: '%v', actual: '%v'", http.StatusBadRequest, statusCode)
		}
	})
}

func TestTodoListEndpointsSuccess(t *testing.T) {
	t.Run("adding a new list object succeeds", func(t *testing.T) {
		setUp()
		//Arrange
		list := database.TodoList{Title: "testTitle", TodoIds: []string{}, ID: "0"}
		var buffer bytes.Buffer
		err := json.NewEncoder(&buffer).Encode(list)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}

		request, err := http.NewRequest(http.MethodPost, "/todoList", &buffer)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}
		response := httptest.NewRecorder()

		//Act
		endpoints.TodoListHandler(response, request)

		//Assert
		statusCode := response.Result().StatusCode
		if http.StatusOK != statusCode {
			t.Errorf("assertion failed. expected: '%v', actual: '%v'", http.StatusOK, statusCode)
		}
		if database.TodoListSlice.Exists(list) < 0 {
			t.Errorf("assertion failed. object was not created properly in database")
		}
	})

	t.Run("updating an existing list object succeeds", func(t *testing.T) {
		setUp()
		//Arrange
		list := database.TodoList{Title: "testTitle", TodoIds: []string{}, ID: "0"}
		var buffer bytes.Buffer
		err := json.NewEncoder(&buffer).Encode(list)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}
		//Add object to database
		request, err := http.NewRequest(http.MethodPost, "/todoList", &buffer)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}
		response := httptest.NewRecorder()
		endpoints.TodoListHandler(response, request)
		//Prepare patch request
		err = json.NewEncoder(&buffer).Encode(list)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}
		request, err = http.NewRequest(http.MethodPatch, "/todoList", &buffer)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}
		response = httptest.NewRecorder()

		//Act
		endpoints.TodoListHandler(response, request)

		//Assert
		statusCode := response.Result().StatusCode
		if http.StatusOK != statusCode {
			t.Errorf("assertion failed. expected: '%v', actual: '%v'", http.StatusOK, statusCode)
		}
		if database.TodoListSlice.Exists(list) != 0 {
			t.Errorf("assertion failed. object was not created properly in database")
		}
	})
}

func TestTodoListEndpointsFailure(t *testing.T) {
	t.Run("adding an existing list object results in a BadRequest error", func(t *testing.T) {
		setUp()
		//Arrange
		list := database.TodoList{Title: "testTitle", TodoIds: nil, ID: "0"}
		var buffer bytes.Buffer
		err := json.NewEncoder(&buffer).Encode(list)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}

		request, err := http.NewRequest(http.MethodPost, "/todoList", &buffer)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}
		response := httptest.NewRecorder()

		//Act
		endpoints.TodoListHandler(response, request)
		//Resend the same request
		request, err = http.NewRequest(http.MethodPost, "/todoList", &buffer)
		if err != nil {
			t.Errorf("failed to execute act block of the test: %v", err)
		}
		endpoints.TodoListHandler(response, request)

		//Assert
		statusCode := response.Result().StatusCode
		if http.StatusBadRequest != statusCode {
			t.Errorf("assertion failed. expected: '%v', actual: '%v'", http.StatusBadRequest, statusCode)
		}
	})

	t.Run("updating a non existing list object results in a BadRequest error", func(t *testing.T) {
		setUp()
		//Arrange
		list := database.TodoList{Title: "testTitle", TodoIds: nil, ID: "0"}
		var buffer bytes.Buffer
		err := json.NewEncoder(&buffer).Encode(list)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}

		request, err := http.NewRequest(http.MethodPatch, "/todoList", &buffer)
		if err != nil {
			t.Errorf("failed to execute arrange block of the test: %v", err)
		}
		response := httptest.NewRecorder()

		//Act
		endpoints.TodoListHandler(response, request)

		//Assert
		statusCode := response.Result().StatusCode
		if http.StatusBadRequest != statusCode {
			t.Errorf("assertion failed. expected: '%v', actual: '%v'", http.StatusBadRequest, statusCode)
		}
	})
}
