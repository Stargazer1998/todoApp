package tests

import (
	"testing"
	services "todoApp/services"
	tododatabase "todoApp/tododb"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

// Define the behavior of the mock repository
func (m *MockRepository) GetTodos() ([]tododatabase.TodoItem, error) {
	args := m.Called()
	return args.Get(0).([]tododatabase.TodoItem), args.Error(1)
}

func (m *MockRepository) CreateTodoItem(title string, completed bool) (tododatabase.TodoItem, error) {
	args := m.Called(title, completed)
	return args.Get(0).(tododatabase.TodoItem), args.Error(1)
}

func (m *MockRepository) UpdateTodoItem(id int, title string, completed bool) (tododatabase.TodoItem, error) {
	args := m.Called(id, title, completed)
	return args.Get(0).(tododatabase.TodoItem), args.Error(1)
}

func (m *MockRepository) DeleteTodoItem(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestService_AddTodo(t *testing.T) {
	mockRepo := new(MockRepository)
	service := services.NewService(mockRepo)

	mockRepo.On("CreateTodoItem", "New Todo", false).Return(tododatabase.TodoItem{
		ID:        1,
		Title:     "New Todo",
		Completed: false,
	}, nil)

	createdTodo, err := service.AddTodo("New Todo", false)

	assert.NoError(t, err)
	assert.Equal(t, "New Todo", createdTodo.Title)
	assert.Equal(t, false, createdTodo.Completed)
	assert.Greater(t, createdTodo.ID, 0)

	mockRepo.AssertExpectations(t)
}

func TestTodoService_UpdateTodo(t *testing.T) {
	mockRepo := new(MockRepository)

	createdTodo := tododatabase.TodoItem{
		ID:        1,
		Title:     "Test Todo",
		Completed: false,
	}

	mockRepo.On("UpdateTodoItem", 1, createdTodo).Return(tododatabase.TodoItem{ID: 1, Title: "Updated Todo", Completed: true}, nil)

	service := services.NewService(mockRepo)

	updatedTodo, err := service.UpdateTodoItem(1, "Updated Todo", true)

	assert.NoError(t, err)
	assert.Equal(t, "Updated Todo", updatedTodo.Title)
	assert.Equal(t, true, updatedTodo.Completed)

	mockRepo.AssertExpectations(t)
}

func TestTodoService_DeleteTodo(t *testing.T) {
	mockRepo := new(MockRepository)

	mockRepo.On("DeleteTodoItem", 1).Return(nil)

	service := services.NewService(mockRepo)

	err := service.DeleteTodoItem(1)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
