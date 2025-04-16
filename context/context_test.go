package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	// Context with value
	// This is used to pass request-scoped values down the call chain
	// The value is not safe for concurrent use
	contextA := context.Background()
	contextB := context.WithValue(contextA, "b", "B") // parent context background
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D") // parent context B
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F") // parent context C

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	// access value to parent
	fmt.Println(contextF.Value("f")) // get value from contextF
	fmt.Println(contextF.Value("c")) // get value from contextC
	fmt.Println(contextF.Value("b")) // nil, different parent
	fmt.Println(contextA.Value("b")) // cannot get value from child
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
				case <- ctx.Done(): // cancel the context
					return
				default:
					destination <- counter
					counter++
					time.Sleep(1 * time.Second) // simulate work
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	// Context with cancel
	// This is used to cancel the context and all its children
	// The cancel function should be called when the context is no longer needed
	// The cancel function should be called to release resources
	// The cancel function should be called to avoid memory leaks

	fmt.Println("Total goroutines", runtime.NumGoroutine()) // 2 goroutine

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent) // create a new context with cancel	

	destination := CreateCounter(ctx)
	
	fmt.Println("Total goroutines", runtime.NumGoroutine()) // 3 goroutine

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 { // cancel the context after 10, but goroutine still running
			break
		}
	}

	cancel() // cancel the context

	time.Sleep(2 * time.Second) // wait for goroutine to finish

	fmt.Println("Total goroutines", runtime.NumGoroutine()) //  3 goroutine
}

func TestContextWithTimeout(t *testing.T) {
	// Context with timeout
	// This is used to set a timeout for the context
	// The timeout is used to cancel the context after a certain period of time
	// The cancel function should be called when the context is no longer needed
	// The cancel function should be called to release resources
	// The cancel function should be called to avoid memory leaks

	fmt.Println("Total goroutines", runtime.NumGoroutine()) // 2 goroutine

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5 * time.Second) // create a new context with timeout	
	defer cancel() // cancel the context when done

	destination := CreateCounter(ctx)

	fmt.Println("Total goroutines", runtime.NumGoroutine()) // 3 goroutine

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second) // wait for goroutine to finish

	fmt.Println("Total goroutines", runtime.NumGoroutine()) //  3 goroutine
}

func TestContextWithDeadline(t *testing.T) {
	// Context with deadline
	// This is used to set a deadline for the context
	// The deadline is used to cancel the context after a certain period of time
	// The cancel function should be called when the context is no longer needed
	// The cancel function should be called to release resources
	// The cancel function should be called to avoid memory leaks

	fmt.Println("Total goroutines", runtime.NumGoroutine()) // 2 goroutine

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5 * time.Second)) // create a new context with deadline	
	defer cancel() // cancel the context when done

	destination := CreateCounter(ctx)

	fmt.Println("Total goroutines", runtime.NumGoroutine()) // 3 goroutine

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second) // wait for goroutine to finish

	fmt.Println("Total goroutines", runtime.NumGoroutine()) //  3 goroutine
}