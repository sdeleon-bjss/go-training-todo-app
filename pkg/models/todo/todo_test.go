package todo_test

import (
	"bjss-todo-app/pkg/models/todo"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	testTodos := []todo.Todo{
		{ID: 1, Task: "Groceries", Status: todo.StatusComplete},
		{ID: 2, Task: "Laundry", Status: todo.StatusInProgress},
		{ID: 3, Task: "Dishes", Status: todo.StatusIncomplete},
		{ID: 4, Task: "Meeting", Status: todo.StatusCancelled},
	}

	err := todo.WriteToFile("test_part_1_todos.json", testTodos...)
	if err != nil {
		t.Errorf("Error writing to file: %v", err)
	}
}
