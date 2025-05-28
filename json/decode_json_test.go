package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	// Example JSON data to decode
	jsonData := `{
		"name": "John Doe",
		"age": 30,
		"address": "123 Main St",
		"age": 30,
		"married": true
	}`

	var data map[string]interface{}

	// Decode the JSON data
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	// Check if the data is decoded correctly
	if data["name"] != "John Doe" || data["age"] != float64(30) || data["address"] != "123 Main St" {
		t.Errorf("Decoded data does not match expected values: %v", data)
	}
	fmt.Printf("Decoded data: %+v\n", data)
}

func TestDecodeJsonToStruct(t *testing.T) {
	// Example JSON data to decode
	jsonData := `{
		"name": "John Doe",
		"age": 30,
		"address": "123 Main St",
		"married": true
	}`

	type Person struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Address string `json:"address"`
		Married bool   `json:"married"`
	}

	var person Person

	// Decode the JSON data into the struct
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		t.Fatalf("Failed to decode JSON to struct: %v", err)
	}

	// Check if the struct is populated correctly
	if person.Name != "John Doe" || person.Age != 30 || person.Address != "123 Main St" || !person.Married {
		t.Errorf("Decoded struct does not match expected values: %+v", person)
	}
	fmt.Printf("Decoded struct: %+v\n", person)
}