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

	println("--- simulating concurrency --- ")
	helloCompleted := make(chan bool)

	go sayHello("hey there", helloCompleted)
	go sayHelloDelay("this is a delayed hello ~", 3, helloCompleted)
	go sayHello("hello there", helloCompleted)
	go sayHelloDelay("this is a another delayed hello ~", 2, helloCompleted)
	go sayHello("hey again", helloCompleted)
	go sayHello("hi!", helloCompleted)

	for hello := range helloCompleted {
		fmt.Println(hello)
	}
}

func sayHello(msg string, completed chan bool) {
	fmt.Println(msg)
	completed <- true
}

func sayHelloDelay(msg string, delay time.Duration, completed chan bool) {
	time.Sleep(delay * time.Second)
	fmt.Println(msg)
	completed <- true
	close(completed)
}
