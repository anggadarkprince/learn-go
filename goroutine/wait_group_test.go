package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsync(group *sync.WaitGroup) {
	defer group.Done() // Decrement the counter when the goroutine completes

	group.Add(1) // Increment the counter before starting the goroutine

	fmt.Println("Hello, Angga!")

	// Simulate some work
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	// Start a goroutine
	for range 10 {
		go RunAsync(group)
	}

	// Wait for the goroutine to finish
	group.Wait()

	fmt.Println("After waiting async")
}