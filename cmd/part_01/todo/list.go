package todo

import (
	"bjss-todo-app/pkg/models"
	"encoding/json"
	"fmt"
	"log"
)

func List(todos ...models.Todo) {
	for _, item := range todos {
		println(item.Task)
	}
}

func ListAsJSON(todos ...models.Todo) {
	todosJSON, err := json.Marshal(todos)
	if err != nil {
		log.Fatalf("Cannot marshal todos: %v", err)
	}

	fmt.Println(string(todosJSON))
}
