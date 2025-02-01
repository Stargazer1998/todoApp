package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	todoApi "todoApp/api"
	tododb "todoApp/tododb"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// let's mock the Service Interface
type MockService struct {
	mock.Mock
}

func (m *MockService) GetAllTodos() ([]tododb.TodoItem, error) {
	args := m.Called()
	return args.Get(0).([]tododb.TodoItem), args.Error(1)
}

func (m *MockService) AddTodo(title string, completed bool) (tododb.TodoItem, error) {
	args := m.Called(title, completed)
	return args.Get(0).(tododb.TodoItem), args.Error(1)
}

func (m *MockService) UpdateTodoItem(id int, title string, completed bool) (tododb.TodoItem, error) {
	args := m.Called(id, title, completed)
	return args.Get(0).(tododb.TodoItem), args.Error(1)
}

func (m *MockService) DeleteTodoItem(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetTodosHandler(t *testing.T) {
	mockService := new(MockService)
	todos := []tododb.TodoItem{{ID: 1, Title: "Test Todo Item", Completed: false}}
	mockService.On("GetAllTodos").Return(todos, nil)

	req, _ := http.NewRequest("GET", "/todos", nil)
	r := httptest.NewRecorder()
	handler := todoApi.GetTodosHandler(mockService)

	handler.ServeHTTP(r, req)

	assert.Equal(t, http.StatusOK, r.Code)
	var response []tododb.TodoItem
	json.Unmarshal(r.Body.Bytes(), &response)
	assert.Equal(t, todos, response)
}

func TestAddTodoItemHandler(t *testing.T) {
	mockService := new(MockService)
	newTodo := tododb.TodoItem{Title: "New Task", Completed: false}
	createdTodo := tododb.TodoItem{ID: 1, Title: "New Task", Completed: false}

	mockService.On("AddTodo", newTodo.Title, newTodo.Completed).Return(createdTodo, nil)

	body, _ := json.Marshal(newTodo)
	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
	r := httptest.NewRecorder()
	handler := todoApi.AddTodoItemHandler(mockService)

	handler.ServeHTTP(r, req)

	assert.Equal(t, http.StatusCreated, r.Code)
	var response tododb.TodoItem
	json.Unmarshal(r.Body.Bytes(), &response)
	assert.Equal(t, createdTodo, response)
}

func TestUpdateTodoItemHandler(t *testing.T) {
	mockService := new(MockService)

	inputTodo := tododb.TodoItem{
		ID:        1,
		Title:     "updated title",
		Completed: true,
	}

	mockService.On("UpdateTodoItem", inputTodo.ID, inputTodo.Title, inputTodo.Completed).Return(inputTodo, nil)

	body, _ := json.Marshal(inputTodo)
	req, err := http.NewRequest("PUT", "/todos", bytes.NewBuffer(body))

	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := todoApi.UpdateTodoItemHandler(mockService)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var updatedTodo tododb.TodoItem

	err = json.NewDecoder(rr.Body).Decode(&updatedTodo)
	assert.NoError(t, err)
	assert.Equal(t, inputTodo, updatedTodo)

	mockService.AssertExpectations(t)

}

func TestDeleteTodoItemHandler(t *testing.T) {
	mockService := new(MockService)
	id := 100
	mockService.On("DeleteTodoItem", id).Return(errors.New("not found"))

	req, _ := http.NewRequest("DELETE", "/todos/100", nil)
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/todos/{id}", todoApi.DeleteTodoItemHandler(mockService))
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Contains(t, rr.Body.String(), "Todo item not found")

}
