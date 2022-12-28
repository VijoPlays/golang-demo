package database

import "errors"

var todos []Todo
var todoLists []TodoList

func AddTodo(todo Todo) error {
	todo.Model.Type = todoType

	exists, err := todo.Model.exists()
	if err != nil {
		return err
	}
	if exists {
		return errors.New("This title is already taken!")
	}

	todos = append(todos, todo)
	return nil
}

func AddTodoList(list TodoList) error {
	list.model.Type = todoListType

	exists, err := list.model.exists()
	if err != nil {
		return err
	}
	if exists {
		return errors.New("This title is already taken!")
	}

	todoLists = append(todoLists, list)
	return nil
}

func UpdateTodo(todo Todo) error {
	todo.Model.Type = todoType

	exists, err := todo.Model.exists()
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Could not find the provided Title!")
	}

	todos = append(todos, todo)
	return nil
}

func UpdateTodoList(list TodoList) error {
	list.model.Type = todoListType

	exists, err := list.model.exists()
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Could not find the provided Title!")
	}

	todoLists = append(todoLists, list)
	return nil
}

// Checks that a model already exists in the database.
func (model DatabaseModel) exists() (bool, error) {
	switch model.Type {

	case todoType:
		for _, v := range todos {
			if v.Model.Title == model.Title {
				return true, nil
			}
		}
		return false, nil

	case todoListType:
		for _, v := range todoLists {
			if v.model.Title == model.Title {
				return true, nil
			}
		}
		return false, nil

	default:
		return false, errors.New("Could not identify type!")
	}
}
