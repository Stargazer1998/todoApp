package services

import (
	"todoApp/tododb"
)

type ServiceInterface interface {
	GetAllTodos() ([]tododb.TodoItem, error)
	UpdateTodoItem(id int, title string, completed bool) (tododb.TodoItem, error)
	AddTodo(title string, completed bool) (tododb.TodoItem, error)
	DeleteTodoItem(id int) error
}
