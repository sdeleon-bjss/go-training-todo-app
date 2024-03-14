package handlers

import (
	"github.com/sdeleon-bjss/pkg/todo"
	"html/template"
	"net/http"
	"os"
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
	// if using air (cd to cmd/part_02/web and run air in terminal) then use these paths
	//pathToTemplates := "templates/*.gohtml"
	//pathToPartials := "templates/partials/*.gohtml"

	// if running from project root
	pathToTemplates := "cmd/part_02/web/templates/*.gohtml"
	pathToPartials := "cmd/part_02/web/templates/partials/*.gohtml"

	if _, err := os.Stat(pathToTemplates); os.IsNotExist(err) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := os.Stat(pathToPartials); os.IsNotExist(err) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseGlob(pathToTemplates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err = tmpl.ParseGlob(pathToPartials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//tmpl, err := template.ParseFiles("cmd/part_02/web/templates/index.gohtml", "cmd/part_02/web/templates/partials/getTodos.gohtml")
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

// Notes:
// net/http allows you to attach data when executing a template
// - you can access this in the template files using {{.Title}} or {{.Message}}
// template.ParseGlob() is used to parse multiple templates at once
// if using air, you have to run program from cmd/part_02/web and not at root
// - this pkg is nice for hot reloading, saves lots of time in dx
