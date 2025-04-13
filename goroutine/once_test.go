package goroutine

import (
	"sync"
	"testing"
	"time"
)

func TestOnce(t *testing.T) {
	// Once is used to ensure that a function is only executed once.
	// It is useful for initializing resources that should only be created once.
	var once sync.Once
	var value int

	// Function to be executed only once
	initialize := func() {
		value++
	}

	// Start multiple goroutines
	for range 10 {
		go func() {
			once.Do(initialize)
		}()
	}

	time.Sleep(1 * time.Second) // Wait for goroutines to finish

	t.Log("Value:", value) // Should print "Value: 1"
}