package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func TodosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// with id param
		if r.URL.Query().Get("id") != "" {
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

			return
		}

		// w/o params
		println("---")
		results := todos.GetAll()

		err := json.NewEncoder(w).Encode(results)
		if err != nil {
			return
		}

		return
	case "POST":

	case "PUT":

	case "DELETE":

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
