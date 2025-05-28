package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestJsonDecoder(t *testing.T) {
	reader, _ := os.Open("customer.json")

	decoder := json.NewDecoder(reader)
	customers := &Customer{}

	decoder.Decode(customers)

	fmt.Printf("Decoded Customers: %+v\n", customers)
}

func TestJsonEncoder(t *testing.T) {
	writer, _ := os.Create("customer_output.json")
	defer writer.Close()

	encoder := json.NewEncoder(writer)
	customers := &Customer{
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       28,
		Married:   false,
		Hobbies:   []string{"painting", "hiking"},
	}

	_ = encoder.Encode(customers)

	fmt.Println(customers)
}