package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)


var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	cond.L.Lock() // Lock the mutex before waiting
	defer cond.L.Unlock() // Unlock the mutex after waiting
	defer group.Done() // Decrement the counter when the goroutine completes

	// Wait for the condition to be met
	cond.Wait()

	fmt.Println("Condition met, value:", value)
}

func TestCond(t *testing.T) {
	// Cond is used to synchronize goroutines.
	// It allows one or more goroutines to wait for a condition to be met.
	// It is useful for implementing producer-consumer patterns.
	for i := range 10 {
		group.Add(1)
		go WaitCondition(i)
	}

	go func ()  {
		for range 10 {
			time.Sleep(1 * time.Second) // Simulate some work
			cond.Signal() // Notify ONE waiting goroutine to run
			//cond.Broadcast() // Notify ALL waiting goroutines	to run
		}
	}();

	group.Wait()
}
