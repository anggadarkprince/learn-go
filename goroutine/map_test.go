package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done() // Decrement the counter when the goroutine completes
	// Add value to the map
	group.Add(1) // Increment the counter before starting the goroutine
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := range 100 {
		go AddToMap(data, i, group)
	}

	group.Wait() // Wait for all goroutines to finish
	data.Range(func(key, value any) bool {
		fmt.Println("Key: ", key, "Value", value)
		return true // Continue iterating
	});
}