package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	services "todoApp/services"
	tododb "todoApp/tododb"

	"github.com/gorilla/mux"
)

func GetTodosHandler(service services.ServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		todos, err := service.GetAllTodos()

		if err != nil {
			http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
			return
		}

		// Marshal Todos into Json
		w.Header().Set("Content-Type", "application/json")
		// take the todo item & encode it to the
		// writer
		err = json.NewEncoder(w).Encode(todos)

		if err != nil {
			http.Error(w, "Error Encoding JSON", http.StatusInternalServerError)
		}

	}
}

func UpdateTodoItemHandler(service services.ServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var newTodoItem tododb.TodoItem

		// take the request body & decode it into the newTodoItem
		err := json.NewDecoder(r.Body).Decode(&newTodoItem)

		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		updatedTodo, err := service.UpdateTodoItem(newTodoItem.ID, newTodoItem.Title, newTodoItem.Completed)

		if err != nil {
			http.Error(w, "Failed to update todo item", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedTodo)

	}
}

func AddTodoItemHandler(service services.ServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var newTodoItem tododb.TodoItem

		err := json.NewDecoder(r.Body).Decode(&newTodoItem)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		createdTodo, err := service.AddTodo(newTodoItem.Title, newTodoItem.Completed)

		if err != nil {
			http.Error(w, "Failed to add todo", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdTodo)

	}
}

func DeleteTodoItemHandler(service services.ServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])

		if err != nil {
			http.Error(w, "Invalid todo ID", http.StatusBadRequest)
			return
		}

		err = service.DeleteTodoItem(id)

		if err != nil {
			if err.Error() == "not found" {
				http.Error(w, "Todo item not found", http.StatusNotFound)
				return
			}

			http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	}
}
