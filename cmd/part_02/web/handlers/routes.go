package handlers

import (
	"bjss-todo-app/pkg/todo"
	"html/template"
	"net/http"
)

var todos = todo.InitializeTodos()

func SetupRoutes() {
	// api
	http.HandleFunc("/api/todos", TodosHandler)

	// page
	http.HandleFunc("/", home)
}

type pageData struct {
	Title   string
	Message string
	Todos   todo.Todos
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("cmd/part_02/web/templates/index.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := pageData{
		Title:   "BJSS Go Academy | Todo App",
		Message: "Hello world test msg",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		return
	}
}
