package tododb

type TodoItem struct {
	ID        int
	Title     string
	Completed bool
}

// type DataBase interface {
// 	CreateTodoItem(item TodoItem) (TodoItem, error)
// 	GetTodoItem(id int) (TodoItem, error)
// 	UpdateTodoItem(id int, item TodoItem) (TodoItem, error)
// 	DeleteTodoItem(id int) error
// }
