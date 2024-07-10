package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Function that simulates a task and writes the result to a file
func task(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the WaitGroup that this goroutine is done

	// Simulate some work with sleep
	time.Sleep(time.Duration(id) * time.Second)
	result := fmt.Sprintf("task %d done\n", id)

	// Create the directory if it doesn't exist
	dir := "./go/async/"
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Printf("Task %d: Error creating directory: %v\n", id, err)
		return
	}

	// Create file with the name text_{id}.txt in the specified directory
	filename := fmt.Sprintf("%stext_%d.txt", dir, id)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Task %d: Error creating file: %v\n", id, err)
		return
	}
	defer file.Close()

	// Write result to file
	if _, err := file.WriteString(result); err != nil {
		fmt.Printf("Task %d: Error writing to file: %v\n", id, err)
	}
}

func main() {
	var wg sync.WaitGroup

	// Launch multiple tasks concurrently
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go task(i, &wg)
	}

	// Wait for all tasks to complete
	wg.Wait()
	fmt.Println("All tasks completed")
}
