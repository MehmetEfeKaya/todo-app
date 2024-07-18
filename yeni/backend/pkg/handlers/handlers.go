package handlers

import (
	"encoding/json"
	"net/http"
	"yeni/backend/pkg/database"
	"yeni/backend/pkg/models"

	"github.com/gorilla/mux"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := mux.NewRouter()

	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	json.NewDecoder(r.Body).Decode(&todo)
	database.CreateTodo(todo)
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	params := mux.Vars(r)

	json.NewDecoder(r.Body).Decode(&todo)
	database.UpdateTodo(params["id"], todo)
	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	database.DeleteTodo(params["id"])
	w.WriteHeader(http.StatusNoContent)
}
