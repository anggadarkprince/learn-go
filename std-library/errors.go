package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError = errors.New("validation error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}
	if id == "0" {
		return NotFoundError
	}
	return nil
}

func main() {
	err := GetById("")
	if errors.Is(err, ValidationError) {
		fmt.Println("validation error")
	} else if errors.Is(err, NotFoundError) {
		fmt.Println("not found error")
	} else {
		fmt.Println("success")
	}
}