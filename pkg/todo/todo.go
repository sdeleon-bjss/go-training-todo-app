package todo

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
)

const (
	StatusInProgress = "In Progress"
	StatusIncomplete = "Incomplete"
	StatusCancelled  = "Cancelled"
	StatusComplete   = "Complete"
)

type Todo struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
}

// --- Part 1 start

func List(todos ...Todo) {
	for _, item := range todos {
		fmt.Println(item)
	}
}

func ListAsJSON(todos ...Todo) {
	todosJSON, err := json.Marshal(todos)
	if err != nil {
		log.Fatalf("Cannot marshal todos: %v", err)
	}

	fmt.Println(string(todosJSON))
}

func WriteToFile(fileName string, todos ...Todo) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(todos); err != nil {
		log.Fatalf("Cannot encode todos: %v", err)
	}

	return nil
}

func ReadFromFile(fileName string) ([]Todo, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var todos []Todo
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&todos); err != nil {
		return nil, err
	}

	return todos, nil
}

// --- Part 1 end

// --- Part 2 start

type Todos struct {
	Todos map[int]Todo
}

func (t Todos) Create(task string) Todo {
	todo := Todo{
		ID:     rand.Intn(10000),
		Task:   task,
		Status: StatusInProgress,
	}

	t.Todos[todo.ID] = todo

	return todo
}

func (t *Todos) Read(id int) (Todo, error) {
	todo, ok := t.Todos[id]
	if !ok {
		return Todo{}, fmt.Errorf("todo with ID %d not found", id)
	}

	return todo, nil
}

func (t *Todos) Update(todoToUpdate Todo) (Todo, error) {
	_, ok := t.Todos[todoToUpdate.ID]
	if !ok {
		return Todo{}, fmt.Errorf("todo with ID %d not found", todoToUpdate.ID)
	}

	t.Todos[todoToUpdate.ID] = todoToUpdate

	return todoToUpdate, nil
}

func (t *Todos) Delete(id int) error {
	_, ok := t.Todos[id]
	if !ok {
		return fmt.Errorf("todo with ID %d not found", id)
	}

	delete(t.Todos, id)

	existingTodos := make([]Todo, 0, len(t.Todos))
	for _, todo := range t.Todos {
		existingTodos = append(existingTodos, todo)
	}

	err := WriteToFile("dummy_todos.json", existingTodos...)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) List() {
	for _, item := range t.Todos {
		fmt.Printf("ID: %d, Task: %s, Status: %s\n", item.ID, item.Task, item.Status)
	}
}

func (t *Todos) GetAll() []Todo {
	var todos []Todo
	for _, item := range t.Todos {
		todos = append(todos, item)
	}

	return todos
}

func (t *Todo) SaveToExistingFile(fileName string) error {
	existingTodos, err := ReadFromFile(fileName)
	if err != nil {
		return err
	}

	existingTodos = append(existingTodos, *t)

	err = WriteToFile(fileName, existingTodos...)
	if err != nil {
		log.Fatalf("Cannot encode todos: %v", err)
	}

	return nil
}

func InitializeTodos() Todos {
	todos := make(map[int]Todo)

	todosFromFile, err := ReadFromFile("dummy_todos.json")
	if err != nil {
		log.Printf("Cannot read from file: %v", err)
	}

	if len(todosFromFile) > 0 {
		for _, todo := range todosFromFile {
			todos[todo.ID] = todo
		}
		return Todos{Todos: todos}
	}

	seedData := []Todo{
		{ID: 1, Task: "Buy groceries", Status: StatusInProgress},
		{ID: 2, Task: "Do laundry", Status: StatusInProgress},
		{ID: 3, Task: "Clean the house", Status: StatusInProgress},
	}

	for _, todo := range seedData {
		todos[todo.ID] = todo
	}

	err = WriteToFile("dummy_todos.json", seedData...)
	if err != nil {
		log.Fatalf("Cannot write to file: %v", err)
	}

	return Todos{Todos: todos}
}

// --- Part 2 end
