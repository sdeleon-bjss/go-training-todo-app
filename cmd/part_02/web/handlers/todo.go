package handlers

import (
	"bjss-todo-app/pkg/todo"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// with id query param
		if r.URL.Query().Get("id") != "" {
			getTodoByID(w, r)
			return
		}

		getTodos(w, r)
		return
	case "POST":
		createTodo(w, r)

	case "PUT":
		updateTodo(w, r)

	case "DELETE":
		deleteTodo(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTodoByID(w http.ResponseWriter, r *http.Request) {
	queryID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(queryID)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, err := todos.Read(id)
	if err != nil {
		return
	}

	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		return
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	results := todos.GetAll()

	err := json.NewEncoder(w).Encode(results)
	if err != nil {
		return
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo todo.Todo

	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newTodo = todos.Create(newTodo.Task)

	err = newTodo.SaveToExistingFile("dummy_todos.json")
	if err != nil {
		return
	}

	err = json.NewEncoder(w).Encode(newTodo)
	if err != nil {
		return
	}
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	var todoToUpdate todo.Todo

	err := json.NewDecoder(r.Body).Decode(&todoToUpdate)
	if err != nil {
		fmt.Println("error decoding: ", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("todo parsed: ", todoToUpdate)

	updatedTodo, err := todos.Update(todoToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(updatedTodo)
	if err != nil {
		return
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	queryID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(queryID)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = todos.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
