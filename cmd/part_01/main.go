package main

import (
	"bjss-todo-app/pkg/models/todo"
	"fmt"
	"time"
)

func main() {
	todos := []todo.Todo{
		{ID: 1, Task: "Groceries", Status: todo.StatusComplete},
		{ID: 2, Task: "Laundry", Status: todo.StatusInProgress},
		{ID: 3, Task: "Dishes", Status: todo.StatusIncomplete},
		{ID: 4, Task: "Meeting", Status: todo.StatusCancelled},
	}

	// p1 - 10
	todo.List(todos...)
	// p1 - 11
	todo.ListAsJSON(todos...)
	// p1 - 12
	err := todo.WriteToFile("part_1_todos.json", todos...)
	if err != nil {
		return
	}
	// p1 - 13
	todosFromFile, err := todo.ReadFromFile("part_1_todos.json")
	if err != nil {
		return
	}
	println("Todos read from file:")
	todo.List(todosFromFile...)

	// p1 - 14
	println("--- simulating concurrency --- ")

	sharedData := 0

	done := make(chan bool)
	num := make(chan int, 20)

	go func() {
		for i := range 10 {
			if i%2 == 0 {
				sharedData = i
				fmt.Printf("Go routine 1: %d\n", sharedData)
			}

			time.Sleep(100 * time.Millisecond)
		}

		done <- true
		num <- sharedData
	}()

	go func() {
		for i := range 10 {
			if i%2 != 0 {
				sharedData = i
				fmt.Printf("Go routine 2: %d\n", sharedData)
			}

			time.Sleep(80 * time.Millisecond)
		}

		done <- true
		num <- sharedData
	}()

	<-done
	<-done
	close(num)
}
