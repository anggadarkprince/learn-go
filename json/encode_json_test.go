package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	// Convert the data to JSON format
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// Log the JSON data
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	// Example data to encode
	data := map[string]interface{}{
		"name":    "John Doe",
		"age":     30,
		"address": "123 Main St",
	}

	// Log the JSON representation of the data
	logJson(data)
	logJson([]string{"apple", "banana", "cherry"})
	logJson("Angga")
	logJson(12345)
	logJson(true)
}