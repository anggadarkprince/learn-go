package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHello() {
	fmt.Println("Hello, Goroutine!")
}

func RunGoRoutine(number int) {
	fmt.Println("Run", number)
}

func TestCreateGoroutine(t *testing.T) {
	go RunHello()
	fmt.Println("Main function")

	time.Sleep(1 * time.Second) // Wait for goroutine to finish
	// Note: In a real-world scenario, you would use sync.WaitGroup or channels to wait for goroutines.
	// Using time.Sleep is not a good practice for synchronization.
	// It is used here only to illustrate the concept of goroutines.
	// In production code, you should use proper synchronization mechanisms.
	// For example, you can use a WaitGroup to wait for multiple goroutines to finish.
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	//     defer wg.Done()
}

func TestManyGoroutine(t *testing.T) {
	for i := range 100000 {
		go RunGoRoutine(i)
	}

	time.Sleep(3 * time.Second) // Wait for goroutines to finish
	// Note: In a real-world scenario, you would use sync.WaitGroup or channels to wait for goroutines.
}