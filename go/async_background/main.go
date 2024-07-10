package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// source to learn https://chatgpt.com/share/4a6f430c-36cf-48a5-8340-131c6f0ac638

// Function that simulates a task and writes the result to a file
func task(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the WaitGroup that this goroutine is done

	// Simulate some work with sleep
	time.Sleep(time.Duration(id+10) * time.Second)
	result := fmt.Sprintf("task %d done\n", id)

	// Create the directory if it doesn't exist
	dir := "./go/async_background/"
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
	sigs := make(chan os.Signal, 1) // Buffered channel to hold one signal
	done := make(chan bool, 1)      // Buffered channel to hold one boolean value

	// Notify the sigs channel on receiving SIGINT or SIGTERM signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine to handle the signals
	go func() {
		sig := <-sigs    // Wait for a signal
		fmt.Println()    // Print a newline
		fmt.Println(sig) // Print the received signal
		done <- true     // Signal that we're done
	}()

	// Launch multiple tasks concurrently
	for i := 1; i <= 3; i++ {
		wg.Add(1)       // Increment the WaitGroup counter
		go task(i, &wg) // Run task in a new goroutine
	}

	// Goroutine to wait for all tasks to complete
	go func() {
		wg.Wait()    // Wait for all tasks to complete
		done <- true // Signal that we're done
	}()

	fmt.Println("Awaiting signal")
	<-done // Wait for a signal or all tasks to complete
	fmt.Println("All tasks completed or terminated")
}
