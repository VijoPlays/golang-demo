package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VijoPlays/golang-demo/src/database"
	"github.com/VijoPlays/golang-demo/src/endpoints"
)

//TODO: Get endpoint for assertion

func setUp() {
	database.CleanUpDatabase()
}

func TestTodoEndpointsSuccess(t *testing.T) {
	setUp()
	t.Run("adding a new todo object succeeds", func(t *testing.T) {
		//Arrange
		todo := database.Todo{Title: "testTitle", Done: true}
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
		if !todo.Exists() {
			t.Errorf("assertion failed. object was not created properly in database")
		}
	})

	setUp()
	t.Run("updating an existing todo object succeeds", func(t *testing.T) {
		//Arrange
		todo := database.Todo{Title: "testTitle", Done: true}
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
		if !todo.Exists() {
			t.Errorf("assertion failed. object was not created properly in database")
		}
	})
}

func TestTodoEndpointsFailure(t *testing.T) {
	setUp()
	t.Run("adding an existing todo object results in a BadRequest error", func(t *testing.T) {
		//Arrange
		todo := database.Todo{Title: "testTitle", Done: true}
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

	setUp()
	t.Run("updating a non existing todo object results in a BadRequest error", func(t *testing.T) {
		//Arrange
		todo := database.Todo{Title: "testTitle", Done: true}
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
	setUp()
	t.Run("adding a new list object succeeds", func(t *testing.T) {
		//Arrange
		list := database.TodoList{Title: "testTitle", Todos: []database.Todo{}}
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
		if !list.Exists() {
			t.Errorf("assertion failed. object was not created properly in database")
		}
	})

	setUp()
	t.Run("updating an existing list object succeeds", func(t *testing.T) {
		//Arrange
		list := database.TodoList{Title: "testTitle", Todos: []database.Todo{}}
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
		if !list.Exists() {
			t.Errorf("assertion failed. object was not created properly in database")
		}
	})
}

func TestTodoListEndpointsFailure(t *testing.T) {
	setUp()
	t.Run("adding an existing list object results in a BadRequest error", func(t *testing.T) {
		//Arrange
		list := database.TodoList{Title: "testTitle", Todos: nil}
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

	setUp()
	t.Run("updating a non existing list object results in a BadRequest error", func(t *testing.T) {
		//Arrange
		list := database.TodoList{Title: "testTitle", Todos: nil}
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
