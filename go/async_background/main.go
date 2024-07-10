package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
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

	// Handle OS signals for graceful shutdown
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// Launch multiple tasks concurrently
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go task(i, &wg)
	}

	// Wait for all tasks to complete
	go func() {
		wg.Wait()
		done <- true
	}()

	fmt.Println("Awaiting signal")
	<-done
	fmt.Println("All tasks completed or terminated")
}
