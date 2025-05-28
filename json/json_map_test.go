package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{
		"name": "John Doe",
		"age": 30,
		"address": "123 Main St"
	}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Printf("Decoded: %+v\n", result)
}

func TestMapEncode(t *testing.T) {
	data := map[string]interface{}{
		"name":    "John Doe",
		"age":     30,
		"address": "123 Main St",
	}

	bytes, _ := json.Marshal(data)

	fmt.Printf("Encoded JSON: %s\n", string(bytes))
}