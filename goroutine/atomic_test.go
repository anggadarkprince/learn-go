package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	// Atomic operations are used to perform operations on variables without using locks.
	// They are useful for performance optimization in concurrent programming.
	// The sync/atomic package provides atomic operations for integers and pointers.
	// Example: atomic.AddInt32, atomic.CompareAndSwapInt32, atomic.LoadInt32, atomic.StoreInt32

	var x int64 = 0
	group := sync.WaitGroup{}

	for range 1000 {
		group.Add(1) // Increment the counter before starting the goroutine
		go func() { // run 1000 goroutines (each goroutine will run in parallel)
			for range 100 {
				//x = x + 1 // each go routine can run run in same time (race condition)
				atomic.AddInt64(&x, 1) // Atomic operation to increment x
			}
			group.Done() // Decrement the counter when the goroutine completes
		}()
	}

	group.Wait() // Wait for all goroutines to finish
	fmt.Println("Final value of x:", x)
}