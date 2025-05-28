package main

import "testing"


type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type Customer struct {
	FirstName string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName string `json:"last_name"`
	Age int `json:"age"`
	Married bool `json:"married"`
	Hobbies []string `json:"hobbies"`
	Addresses []Address `json:"addresses"`
}

func TestJsonObject(t *testing.T) {
	// Create a new Customer object
	customer := Customer{
		FirstName: "John",
		MiddleName: "A.",
		LastName: "Doe",
		Age: 30,
		Married: true,
	}

	// Log the JSON representation of the customer object
	logJson(customer)
}