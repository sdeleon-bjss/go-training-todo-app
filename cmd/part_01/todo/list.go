package todo

import "bjss-todo-app/pkg/models"

func List(todos ...models.Todo) {
	for _, item := range todos {
		println(item.Task)
	}
}
