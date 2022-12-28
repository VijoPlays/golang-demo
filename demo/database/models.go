package database

var todoType = "todo"
var todoListType = "todoList"

// Base model for database.
type DatabaseModel struct {
	Title string `json:"title"`
	Type  string
}

// Object to simulate tasks that are left to do.
//
// Title must stay unique!
type Todo struct {
	Model DatabaseModel
	Done  bool `json:"done"`
}

// Object to simulate multiple tasks grouped in a list that are left to do.
//
// Title must stay unique!
type TodoList struct {
	model DatabaseModel
	Todos []Todo
}
