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

// crud tests
func TestCreate(t *testing.T) {
	todos := todo.Todos{
		Todos: make(map[int]todo.Todo),
	}

	todoCreated := todos.Create("Groceries")

	if _, ok := todos.Todos[todoCreated.ID]; !ok {
		t.Errorf("Todo not created")
	}
}

func TestRead(t *testing.T) {
	todos := todo.Todos{
		Todos: make(map[int]todo.Todo),
	}

	todoCreated := todos.Create("Groceries")

	todoRead, err := todos.Read(todoCreated.ID)
	if err != nil {
		t.Errorf("Error reading todo: %v", err)
	}

	if todoRead.ID != todoCreated.ID {
		t.Errorf("Todo ID mismatch")
	}
}

func TestUpdate(t *testing.T) {
	todos := todo.Todos{
		Todos: make(map[int]todo.Todo),
	}

	todoCreated := todos.Create("Groceries")

	todoUpdated, err := todos.Update(todoCreated.ID, "Groceries", todo.StatusComplete)
	if err != nil {
		t.Errorf("Error updating todo: %v", err)
	}

	if todoUpdated.Status != todo.StatusComplete {
		t.Errorf("Todo status not updated")
	}
}

func TestDelete(t *testing.T) {
	todos := todo.Todos{
		Todos: make(map[int]todo.Todo),
	}

	todoCreated := todos.Create("Groceries")

	err := todos.Delete(todoCreated.ID)
	if err != nil {
		t.Errorf("Error deleting todo: %v", err)
	}

	if _, ok := todos.Todos[todoCreated.ID]; ok {
		t.Errorf("Todo not deleted")
	}
}
