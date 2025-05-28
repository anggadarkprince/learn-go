package main

import (
	"encoding/json"
	"testing"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		ID:          1,
		Name:        "Laptop",
		Description: "A high-performance laptop",
	}

	// Log the JSON representation of the product object
	logJson(product)
}

func TestJsonTagDecode(t *testing.T) {
	jsonData := `{"id":1,"name":"Laptop","description":"A high-performance laptop"}`
	jsonBytes := []byte(jsonData)

	product := Product{}
	json.Unmarshal(jsonBytes, &product)
	
	t.Logf("Decoded Product: %+v\n", product)
}