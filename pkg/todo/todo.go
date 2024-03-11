package todo

import (
	"bjss-todo-app/cmd/part_02/web/database"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
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

type Todos struct {
	Todos map[int]Todo
	mu    sync.RWMutex
}

func (t *Todos) Create(task string) (Todo, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	sql := "INSERT INTO todos (task, status, user_id) VALUES ($1, $2, $3) RETURNING id"

	createdTodo := Todo{}
	status := StatusInProgress

	err := database.DB.QueryRow(sql, task, status, 1).Scan(&createdTodo.ID)
	if err != nil {
		log.Fatalf("Error creating todo: %v", err)
		return Todo{}, err
	}

	createdTodo.Task = task
	createdTodo.Status = status

	return createdTodo, nil
}

func (t *Todos) Read(id int) (Todo, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	sql := "SELECT id, task, status FROM todos WHERE id = $1"

	todo := Todo{}
	err := database.DB.QueryRow(sql, id).Scan(&todo.ID, &todo.Task, &todo.Status)
	if err != nil {
		log.Fatalf("Error reading todo: %v", err)
		return Todo{}, err
	}

	return todo, nil
}

func (t *Todos) Update(todoToUpdate Todo) (Todo, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	sql := "UPDATE todos SET task = $1, status = $2 WHERE id = $3 RETURNING id, task, status"

	updatedTodo := Todo{}
	err := database.DB.QueryRow(sql, todoToUpdate.Task, todoToUpdate.Status, todoToUpdate.ID).Scan(&updatedTodo.ID, &updatedTodo.Task, &updatedTodo.Status)
	if err != nil {
		log.Fatalf("Error updating todo: %v", err)
		return Todo{}, err
	}

	return updatedTodo, nil
}

func (t *Todos) Delete(id int) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	sql := "DELETE FROM todos WHERE id = $1"

	_, err := database.DB.Exec(sql, id)
	if err != nil {
		log.Fatalf("Error deleting todo: %v", err)
		return err
	}

	return nil
}

func (t *Todos) List() {
	sql := "SELECT id, task, status FROM todos"

	rows, err := database.DB.Query(sql)
	if err != nil {
		log.Fatalf("Error getting todos: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var readTodo Todo

		err := rows.Scan(&readTodo.ID, &readTodo.Task, &readTodo.Status)
		if err != nil {
			log.Fatalf("Error scanning todo: %v", err)
		}

		fmt.Printf("ID: %d, Task: %s, Status: %s\n", readTodo.ID, readTodo.Task, readTodo.Status)
	}
}

func (t *Todos) GetAll() ([]Todo, error) {
	var todosList []Todo

	sql := "SELECT id, task, status FROM todos"

	rows, err := database.DB.Query(sql)
	if err != nil {
		log.Fatalf("Error getting todosList: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var readTodo Todo

		err := rows.Scan(&readTodo.ID, &readTodo.Task, &readTodo.Status)
		if err != nil {
			log.Fatalf("Error scanning todo: %v", err)
			return nil, err
		}

		todosList = append(todosList, readTodo)
	}

	return todosList, nil

}

// keeping around for other programs using it still
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
