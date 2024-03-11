package handlers

import (
	"bjss-todo-app/pkg/todo"
	"html/template"
	"net/http"
)

var todos = todo.Todos{}

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
	// if using air:
	//pathToTemplates := "templates/*.gohtml"
	//pathToPartials := "templates/partials/*.gohtml"
	//
	//if _, err := os.Stat(pathToTemplates); os.IsNotExist(err) {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//if _, err := os.Stat(pathToPartials); os.IsNotExist(err) {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//tmpl, err := template.ParseGlob(pathToTemplates)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//tmpl, err = tmpl.ParseGlob(pathToPartials)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	tmpl, err := template.ParseFiles("cmd/part_02/web/templates/index.gohtml", "cmd/part_02/web/templates/partials/getTodos.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	allTodos, err := todos.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := pageData{
		Title:   "BJSS Go Academy | Todo App",
		Message: "Hello world from the server",
		Todos:   allTodos,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		return
	}
}
