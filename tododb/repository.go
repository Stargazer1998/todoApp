package tododb

// interface: interactions w/ Database
type Repository interface {
	GetTodos() ([]TodoItem, error)
	CreateTodoItem(title string, completed bool) (TodoItem, error)
	UpdateTodoItem(id int, title string, completed bool) (TodoItem, error)
	DeleteTodoItem(id int) error
}
