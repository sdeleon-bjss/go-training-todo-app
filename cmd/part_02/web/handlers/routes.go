package handlers

import (
	"bjss-todo-app/pkg/todo"
	"html/template"
	"net/http"
)

var todos = todo.InitializeTodos()

func SetupRoutes() {
	// api
	http.HandleFunc("/api/todos", todosHandler)

	// page
	http.HandleFunc("/", home)
}

type pageData struct {
	Title   string
	Message string
	Todos   []todo.Todo
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("cmd/part_02/web/templates/index.gohtml", "cmd/part_02/web/templates/partials/getTodos.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := pageData{
		Title:   "BJSS Go Academy | Todo App",
		Message: "Hello world from the server",
		Todos:   todos.GetAll(),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		return
	}
}
