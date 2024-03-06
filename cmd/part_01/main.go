package main

import (
	"bjss-todo-app/pkg/models/todo"
)

func main() {
	todos := []todo.Todo{
		{ID: 1, Task: "Groceries", Status: todo.StatusComplete},
		{ID: 2, Task: "Laundry", Status: todo.StatusInProgress},
		{ID: 3, Task: "Dishes", Status: todo.StatusIncomplete},
		{ID: 4, Task: "Meeting", Status: todo.StatusCancelled},
	}

	todo.List(todos...)
	todo.ListAsJSON(todos...)
	err := todo.WriteToFile("part_1_todos.json", todos...)
	if err != nil {
		return
	}
}
