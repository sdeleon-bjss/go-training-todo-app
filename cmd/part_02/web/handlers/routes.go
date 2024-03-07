package handlers

import (
	"html/template"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/", home)
}

type pageData struct {
	Title   string
	Message string
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
