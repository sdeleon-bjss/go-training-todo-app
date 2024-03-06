package main

import (
	"bjss-todo-app/pkg/models/todo"
	"fmt"
	"time"
)

var sharedNumber int

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
	oddNumbers := []int{1, 3, 5, 7, 9, 11}
	evenNumbers := []int{2, 4, 6, 8, 10}

	numbersUpdatedChan := make(chan int)

	go updateCount(oddNumbers, numbersUpdatedChan)
	go updateCount(evenNumbers, numbersUpdatedChan)

	for updatedNum := range numbersUpdatedChan {
		fmt.Println(updatedNum)
	}

}

func updateCount(numbers []int, numbersUpdatedChan chan int) {
	for _, num := range numbers {
		sharedNumber = num
		fmt.Printf("updating count with %d", num)
		time.Sleep(200 * time.Millisecond)
		numbersUpdatedChan <- sharedNumber
		//close(numbersUpdatedChan)
	}
}
