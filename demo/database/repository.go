package database

import "errors"

var todos []Todo
var todoLists []TodoList

func AddTodo(todo Todo) error {
	exists := todo.Exists()
	if exists {
		return errors.New("This title is already taken!")
	}

	todos = append(todos, todo)
	return nil
}

func AddTodoList(list TodoList) error {
	exists := list.Exists()
	if exists {
		return errors.New("This title is already taken!")
	}

	appendList(list)
	return nil
}

func UpdateTodo(todo Todo) error {
	exists := todo.Exists()
	if !exists {
		return errors.New("Could not find the provided Title!")
	}

	todos = append(todos, todo)
	return nil
}

func UpdateTodoList(list TodoList) error {
	exists := list.Exists()
	if !exists {
		return errors.New("Could not find the provided Title!")
	}

	appendList(list)
	return nil
}

// This is mainly used to make testing easier.
//
// WARNING: Do not pollute your code for testing purposes! This is not good practice and only done in this case to save time on creating new methods to remove Todo/TodoList from the repository.
func CleanUpDatabase() {
	todos = make([]Todo, 0)
	todoLists = make([]TodoList, 0)
}

// Adds the list to the database, and adds any missing todos that are a part of the list to the database as well.
func appendList(list TodoList) {
	for _, v := range list.Todos {
		if !v.Exists() {
			todos = append(todos, v)
		}
	}

	todoLists = append(todoLists, list)
}

// Returns true if an object with this Title already exists in the database.
//
// WARNING: Should be private, was made public to make assertion (for tests) simpler - but shouldn't be done in production code (for that we'd have a GetByTitle method anyway we can use for tests).
func (todo Todo) Exists() bool {
	for _, v := range todos {
		if v.Title == todo.Title {
			return true
		}
	}
	return false
}

// Returns true if an object with this Title already exists in the database.
//
// WARNING: Should be private, was made public to make assertion (for tests) simpler - but shouldn't be done in production code (for that we'd have a GetByTitle method anyway we can use for tests).
func (list TodoList) Exists() bool {
	for _, v := range todoLists {
		if v.Title == list.Title {
			return true
		}
	}
	return false
}
