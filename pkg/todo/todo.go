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

func (t *Todos) Update(id int, task string, status string) (Todo, error) {
	todo, ok := t.Todos[id]
	if !ok {
		return Todo{}, fmt.Errorf("todo with ID %d not found", id)
	}

	todo.Task = task
	todo.Status = status

	t.Todos[id] = todo

	return todo, nil
}

func (t *Todos) Delete(id int) error {
	_, ok := t.Todos[id]
	if !ok {
		return fmt.Errorf("todo with ID %d not found", id)
	}

	delete(t.Todos, id)

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

func InitializeTodos() Todos {
	todos := Todos{
		Todos: make(map[int]Todo),
	}

	dummyTodos, err := ReadFromFile("dummy_todos.json")
	if err != nil {
		return todos
	}

	for _, item := range dummyTodos {
		todos.Todos[item.ID] = item
	}

	return todos
}

// --- Part 2 end
