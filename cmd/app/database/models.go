package database

// Object to simulate tasks that are left to do.
//
// Title must stay unique!
type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// Object to simulate multiple tasks grouped in a list that are left to do.
//
// Title must stay unique!
type TodoList struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	TodoIds []string `json:"todo"`
}
