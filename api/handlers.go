package api

import (
	services "todoApp/services"

	"github.com/gorilla/mux"
)

func InitializeRoutes(todoService services.ServiceInterface) *mux.Router {
	router := mux.NewRouter()

	// TODO: Define API routes and link them to their handlers

	router.HandleFunc("/todos", GetTodosHandler(todoService)).Methods("GET")
	router.HandleFunc("/todos", AddTodoItemHandler(todoService)).Methods("POST")
	router.HandleFunc("/todos/{id:[0-9]+}", UpdateTodoItemHandler(todoService)).Methods("PUT")
	router.HandleFunc("/todos/{id:[0-9]+}", DeleteTodoItemHandler(todoService)).Methods("DELETE")
	return router
}
