package main

import (
	"bjss-todo-app/cmd/part_01/todo"
	"bjss-todo-app/pkg/models"
)

func main() {
	todos := []models.Todo{
		{ID: 1, Task: "Groceries", Status: models.StatusComplete},
		{ID: 2, Task: "Laundry", Status: models.StatusInProgress},
		{ID: 3, Task: "Dishes", Status: models.StatusIncomplete},
		{ID: 4, Task: "Meeting", Status: models.StatusCancelled},
	}

	todo.List(todos...)
	todo.ListAsJSON(todos...)
}
