package main

import (
	"errors"
	"fmt"
	"strconv"
)

type ValidationError struct {
	Message string
}

func (v ValidationError) Error() string {
	return "Invalid: " + v.Message;
}

type NotFoundError struct {
	Message string
}

func (v NotFoundError) Error() string {
	return "Not Found: " + v.Message;
}

func Divider(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("Cannot divide with 0")
	}
	return value / divider, nil
}

func StoreUser(id int, data any) error {
	if id == 0 {
		return &ValidationError{"ID is empty"}
	}
	if id < 0 {
		return &NotFoundError{"ID" + strconv.Itoa(id) + " is not found"}
	}
	return nil
}

func main() {
	result, errorDivider := Divider(10, 0)
	if (errorDivider == nil) {
		fmt.Println(result)
	} else {
		fmt.Println("Error:", errorDivider.Error())
	}

	created := StoreUser(0, "Angga")
	if (created != nil) {
		if validation, ok := created.(*ValidationError); ok {
			fmt.Println("Validation", validation.Error())
		} else if notFound, ok := created.(*NotFoundError); ok {
			fmt.Println("Not found", notFound.Error())
		} else {
			fmt.Println("Unknown error", created.Error())
		}
	} else {
		fmt.Println("Store user success")
	}
}