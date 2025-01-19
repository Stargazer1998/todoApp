package services

import (
	tododb "todoApp/tododb"
)

type Service struct {
	repo tododb.Repository
}

func NewService(repo tododb.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllTodos() ([]tododb.TodoItem, error) {
	return s.repo.GetTodos()
}

func (s *Service) AddTodo(title string, completed bool) (tododb.TodoItem, error) {
	return s.repo.CreateTodoItem(title, completed)
}

func (s *Service) UpdateTodoItem(id int, title string, completed bool) (tododb.TodoItem, error) {
	return s.repo.UpdateTodoItem(id, title, completed)
}

func (s *Service) DeleteTodoItem(id int) error {
	return s.repo.DeleteTodoItem(id)
}
