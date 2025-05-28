package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName: "John",
		MiddleName: "A.",
		LastName: "Doe",
		Age: 30,
		Married: true,
		Hobbies: []string{"reading", "traveling", "coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonData := `{"FirstName":"John","MiddleName":"A.","LastName":"Doe","Age":30,"Married":true,"Hobbies":["reading","traveling","coding"]}`
	jsonBytes := []byte(jsonData)

	customer := Customer{}
	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)	
	}
	fmt.Printf("Decoded Customer: %+v\n", customer)
}


func TestJsonArrayNested(t *testing.T) {
	customer := Customer{
		FirstName: "John",
		Addresses: []Address{
			{Street: "123 Main St", City: "New York", Country: "USA"},
			{Street: "456 Elm St", City: "Los Angeles", Country: "USA"},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayNestedDecode(t *testing.T) {
	jsonData := `{"FirstName":"John","Addresses":[{"Street":"123 Main St","City":"New York","Country":"USA"},{"Street":"456 Elm St","City":"Los Angeles","Country":"USA"}]}`
	jsonBytes := []byte(jsonData)

	customer := Customer{}
	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)	
	}
	fmt.Printf("Decoded Customer with Addresses: %+v\n", customer)
}

func TestJsonOnlyArray(t *testing.T) {
	addresses := []Address{
		{Street: "123 Main St", City: "New York", Country: "USA"},
		{Street: "456 Elm St", City: "Los Angeles", Country: "USA"},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}

func TestJsonOnlyArrayDecode(t *testing.T) {
	jsonData := `[{"Street":"123 Main St","City":"New York","Country":"USA"},{"Street":"456 Elm St","City":"Los Angeles","Country":"USA"}]`
	jsonBytes := []byte(jsonData)

	addresses := []Address{}
	err := json.Unmarshal(jsonBytes, &addresses)
	if err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)	
	}
	fmt.Printf("Decoded Addresses: %+v\n", addresses)
}