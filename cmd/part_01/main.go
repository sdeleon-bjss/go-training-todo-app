package main

import (
	"bjss-todo-app/pkg/models/todo"
)

func main() {
	todos := []todo.Todo{
		{ID: 1, Task: "Groceries", Status: todo.StatusComplete},
		{ID: 2, Task: "Laundry", Status: todo.StatusInProgress},
		{ID: 3, Task: "Dishes", Status: todo.StatusIncomplete},
		{ID: 4, Task: "Meeting 1", Status: todo.StatusCancelled},
		{ID: 5, Task: "Meeting 2", Status: todo.StatusCancelled},
		{ID: 6, Task: "Meeting 3", Status: todo.StatusCancelled},
		{ID: 7, Task: "Exercise", Status: todo.StatusIncomplete},
		{ID: 8, Task: "Shower", Status: todo.StatusIncomplete},
		{ID: 9, Task: "Feed cats", Status: todo.StatusComplete},
		{ID: 10, Task: "Cook dinner", Status: todo.StatusInProgress},
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

	// p1 - 14
	//sharedData := 0
	//
	//done := make(chan bool)
	//num := make(chan int, 20)
	//
	//go func() {
	//	for i := range 10 {
	//		if i%2 == 0 {
	//			sharedData = i
	//			fmt.Printf("Go routine 1: %d\n", sharedData)
	//		}
	//
	//		time.Sleep(100 * time.Millisecond)
	//	}
	//
	//	done <- true
	//	num <- sharedData
	//}()
	//
	//go func() {
	//	for i := range 10 {
	//		if i%2 != 0 {
	//			sharedData = i
	//			fmt.Printf("Go routine 2: %d\n", sharedData)
	//		}
	//
	//		time.Sleep(80 * time.Millisecond)
	//	}
	//
	//	done <- true
	//	num <- sharedData
	//}()
	//
	//<-done
	//<-done
	//close(num)

	// p1 - 15
	//sharedData := 0
	//
	//var mu sync.Mutex
	//var wg sync.WaitGroup
	//
	//wg.Add(2)
	//
	//go func() {
	//	defer wg.Done()
	//
	//	for i := range 10 {
	//		if i%2 == 0 {
	//			mu.Lock()
	//			sharedData = i
	//			mu.Unlock()
	//			fmt.Println("Go routine 1: ", sharedData)
	//		}
	//
	//		time.Sleep(1 * time.Millisecond)
	//	}
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//
	//	for i := range 10 {
	//		if i%2 != 0 {
	//			mu.Lock()
	//			sharedData = i
	//			mu.Unlock()
	//			fmt.Println("Go routine 2: ", sharedData)
	//		}
	//
	//		time.Sleep(1 * time.Millisecond)
	//	}
	//}()
	//
	//wg.Wait()
	//fmt.Println("Final value:", sharedData)

	// p1 - 16
	//var mu sync.Mutex
	//var wg sync.WaitGroup
	//
	//wg.Add(2)
	//
	//go func() {
	//	defer wg.Done()
	//
	//	for _, t := range todos {
	//		mu.Lock()
	//		fmt.Printf("To Do Item:(%d) %s\n", t.ID, t.Task)
	//		mu.Unlock()
	//		time.Sleep(100 * time.Millisecond)
	//	}
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//
	//	for _, t := range todos {
	//		mu.Lock()
	//		fmt.Printf("To DO Status:(%d) %s\n", t.ID, t.Status)
	//		mu.Unlock()
	//		time.Sleep(100 * time.Millisecond)
	//	}
	//}()
	//
	//wg.Wait()
}
