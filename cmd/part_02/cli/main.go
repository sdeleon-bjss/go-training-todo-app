package main

import (
	"bjss-todo-app/pkg/todo"
	"flag"
	"fmt"
)

// TODO - figure out how to keep program open after running a command and not exiting

func main() {
	todos := todo.InitializeTodos()

	// flags
	operation := flag.String("operation", "", "Choose an operation: list, create, read, update or delete")
	createTask := flag.String("new-task", "", "Task for Todo on create mode")
	taskId := flag.Int("id", 0, "ID of the task")
	updateTask := flag.String("task", "", "Updated task description")
	updateStatus := flag.String("status", "", "Updated task status")
	help := flag.String("help", "", "description: Prints this help")

	flag.Parse()

	switch *operation {
	case "help":
		if help == nil || *help != "" {
			flag.PrintDefaults()
		}
	case "list":
		todos.ListInMemory()
	case "create":
		if *createTask == "" {
			println("Task is required")
			return
		}
		created, err := todos.CreateInMemory(*createTask)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Todo successfully created:", created)
	case "read":
		if *taskId <= 0 {
			println("the `id` flag is required for reading")
			return
		}

		foundTodo, err := todos.ReadInMemory(*taskId)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Todo found:", foundTodo)
	case "update":
		if *updateTask == "" {
			println("the `task` flag is required for updating")
			return
		}
		if *updateStatus == "" {
			println("the `status` flag is required for updating")
			return
		}
		if *taskId <= 0 {
			println("the `id` flag is required for updating")
		}

		updatedTodo, err := todos.UpdateInMemory(
			todo.Todo{
				ID:     *taskId,
				Task:   *updateTask,
				Status: *updateStatus,
			})
		if err != nil {
			return
		}
		fmt.Println("Todo successfully updated:", updatedTodo)
	case "delete":
		if *taskId <= 0 {
			println("the `id` flag is required for deleting")
			return
		}

		err := todos.DeleteInMemory(*taskId)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Todo successfully deleted")
	default:
		println("Operation not found, please choose one of the following: list, create, read, update or delete")
	}
}
