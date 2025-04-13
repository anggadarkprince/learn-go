package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// Start a goroutine to send data to the channel
	go func() {
		time.Sleep(2 * time.Second) // Simulate some work
		channel <- "Hello, Angga!"
		fmt.Println("Data sent to channel")
	}()

	// Receive data from the channel
	//t.Log(<- channel)
	message := <-channel // if there is 2 step to get value in channel, the it waits forever

	// Print the received message
	t.Log(message)
}

func GiveMeResponse(channel chan string) { // no need mark as pointer (in and out)
	// Simulate some work
	time.Sleep(2 * time.Second)
	channel <- "Hello, Angga!"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// Start a goroutine to send data to the channel
	go GiveMeResponse(channel)

	// Receive data from the channel
	message := <-channel

	// Print the received message
	t.Log(message)
}

func OnlyIn(channel chan<- string) {
	// Send data to the channel
	channel <- "Hello, Angga!"
}

func OnlyOut(channel <-chan string) {
	// Receive data from the channel
	message := <-channel
	fmt.Println(message)
}

func TestChannelInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// Start a goroutine to send data to the channel
	go OnlyIn(channel)

	// Start a goroutine to receive data from the channel
	go OnlyOut(channel)

	// Wait for a while to allow goroutines to finish
	time.Sleep(2 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 2) // Buffered channel with capacity of 2
	defer close(channel)

	fmt.Println("capacity", cap(channel)) // Print the capacity of the channel
	fmt.Println("current length", len(channel)) // Print the length of data of the channel

	// Send data to the buffered channel
	channel <- "Hello, Angga!"
	fmt.Println("current length", len(channel)) // Print the length of data of the channel
	channel <- "Hello, Ari!"
	channel <- "Hello, Wijaya!" // wait

	// Receive data from the buffered channel
	message1 := <-channel
	message2 := <-channel

	// Print the received messages
	t.Log(message1)
	t.Log(message2)
}

func TestChannelTimeout(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// Start a goroutine to send data to the channel
	go func() {
		time.Sleep(2 * time.Second) // Simulate some work
		channel <- "Hello, Angga!"
	}()

	select {
	case message := <-channel:
		t.Log(message)
	case <-time.After(1 * time.Second): // will timeout since cahnnel is not ready
		t.Log("Timeout waiting for message")
	}
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	// Start a goroutine to send data to the channel
	go func() {
		for i := range 10 {
			channel <- fmt.Sprintf("Message %d", i)
		}
		close(channel)
	}()

	// Use range to receive data from the channel
	for message := range channel {
		t.Log(message)
	}
	fmt.Println("Channel finished")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	// Start a goroutine to send data to channel1
	go func() {
		time.Sleep(2 * time.Second) // Simulate some work
		channel1 <- "Hello from channel 1"
	}()

	// Start a goroutine to send data to channel2
	go func() {
		time.Sleep(3 * time.Second) // Simulate some work
		channel2 <- "Hello from channel 2"
	}()

	// fetch the fastest channel
	// Use select to receive data from either channel
	counter := 0

	for counter < 2 {
		select {
		case message := <-channel1:
			t.Log(message)
			counter++
		case message := <-channel2:
			t.Log(message)
			counter++
		case <-time.After(4 * time.Second):
			t.Log("Timeout waiting for messages")
		}
	}
}

func TestSelectChannelDefault(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	// Start a goroutine to send data to channel1
	go func() {
		time.Sleep(2 * time.Second) // Simulate some work
		channel1 <- "Hello from channel 1"
	}()

	// Start a goroutine to send data to channel2
	go func() {
		time.Sleep(3 * time.Second) // Simulate some work
		channel2 <- "Hello from channel 2"
	}()

	// Use select with default case
	for range 5 {
		select {
		case message := <-channel1:
			t.Log(message)
		case message := <-channel2:
			t.Log(message)
		default:
			t.Log("No messages received, doing other work...")
			time.Sleep(1 * time.Second) // Simulate doing other work
		}
	}
}
