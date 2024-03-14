package main

import (
	"github.com/sdeleon-bjss/cmd/part_02/web/database"
	"github.com/sdeleon-bjss/cmd/part_02/web/handlers"
	"net/http"
)

func main() {
	database.DB = database.Connect()
	defer database.DB.Close()

	handlers.SetupRoutes()

	println("Server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
