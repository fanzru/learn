package main

import (
	"fmt"
	"time"
)

// Worker function that does some work and sends result to channel
func worker(id int, results chan<- string) {
	// Simulate some work with sleep
	time.Sleep(time.Duration(4) * time.Second)
	result := fmt.Sprintf("Worker %d done", id)
	results <- result
}

func main() {
	// Create a channel to communicate results
	results := make(chan string, 3)

	// Launch multiple workers concurrently
	for i := 1; i <= 3; i++ {
		go worker(i, results)
	}

	// Collect results from workers
	for i := 1; i <= 3; i++ {
		fmt.Println(<-results)
	}
}
