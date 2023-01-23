// The database package provides the API with access to an inMemoryDatabase.
package database

import (
	"errors"
	"strconv"
)

// todos is a goType, consisting of a slice of Todo and providing methods for it.
type todos []Todo

// todoLists is a goType, consisting of a slice of TodoList and providing methods for it.
type todoLists []TodoList

// WARNING: Should be private: Made it public to make validation/testing easier, but should not be done in production code.
var TodoSlice todos

// WARNING: Should be private: Made it public to make validation/testing easier, but should not be done in production code.
var TodoListSlice todoLists

//TODO: Fix Postman collection

// AddTodo overwrites the provided ID and then stores the Todo in the database.
// The provided data is not being validated thus far.
func AddTodo(todo Todo) {
	todo.ID = strconv.Itoa(len(TodoSlice))
	TodoSlice = append(TodoSlice, todo)
}

// AddTodoList overwrites the provided ID and then stores the TodoList in the database.
// The provided data is not being validated thus far.
func AddTodoList(list TodoList) {
	list.ID = strconv.Itoa(len(TodoListSlice))
	TodoListSlice = append(TodoListSlice, list)
}

// UpdateTodo checks if the provided todo exists in the database and overwrites it, if it does. Else it'll throw an error.
// The provided data is not being validated thus far.
func UpdateTodo(todo Todo) error {
	index := TodoSlice.Exists(todo)
	if index < 0 {
		return errors.New("could not find the provided item")
	}

	TodoSlice[index] = todo
	return nil
}

// UpdateTodoList checks if the provided list exists in the database and overwrites it, if it does. Else it'll throw an error.
// The provided data is not being validated thus far.
func UpdateTodoList(list TodoList) error {
	index := TodoListSlice.Exists(list)
	if index < 0 {
		return errors.New("could not find the provided item")
	}

	TodoListSlice[index] = list
	return nil
}

// This is mainly used to make testing easier.
//
// WARNING: Do not pollute your code for testing purposes! This is not good practice and only done in this case to save time on creating new methods to remove Todo/TodoList from the repository.
func CleanUpDatabase() {
	TodoSlice = make(todos, 0)
	TodoListSlice = make(todoLists, 0)
}

// Returns the index of the item to replace it easily for updates.
// Returns -1 if the element is not in the list.
//
// WARNING: Should be private, was made public to make assertion (for tests) simpler - but shouldn't be done in production code (for that we'd have a GetByTitle method anyway we can use for tests).
func (t todos) Exists(todo Todo) int {
	for k, v := range t {
		if v.ID == todo.ID {
			return k
		}
	}
	return -1
}

// Returns the index of the item to replace it easily for updates.
// Returns -1 if the element is not in the list.
//
// WARNING: Should be private, was made public to make assertion (for tests) simpler - but shouldn't be done in production code (for that we'd have a GetByTitle method anyway we can use for tests).
func (t todoLists) Exists(list TodoList) int {
	for k, v := range t {
		if v.ID == list.ID {
			return k
		}
	}
	return -1
}
