package main

import (
	"bjss-todo-app/cmd/part_02/web/handlers"
	"net/http"
)

func main() {
	handlers.SetupRoutes()

	println("Server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
