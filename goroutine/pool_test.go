package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	// Pool is a goroutine pool that limits the number of concurrent goroutines.
	// It is useful for controlling resource usage and preventing too many goroutines from running at once.

	// Create a new pool with a maximum of 5 goroutines
	pool := sync.Pool{
		New: func () any  { // Add default value for the pool (if no data is available or all iitems being used)
			return "No name"
		},
	}

	pool.Put("Angga")
	pool.Put("Ari")
	pool.Put("Wijaya")

	// Add tasks to the pool
	for range 10 {
		go func() {
			data := pool.Get()
			defer pool.Put(data) // Return the data to the pool after use

			time.Sleep(1 * time.Second) // Simulate data is not returned yet
			fmt.Println(data)
		}()
	}

	time.Sleep(1 * time.Second) // Wait for goroutines to finish
	fmt.Println("All tasks completed")
}