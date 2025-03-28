package main

import (
	"fmt"
	"reflect"
)

type UserAccount struct {
	Name string `required:"true"`
	Address string `required:"true"`
}

type UserAddress struct {
	Address string
	Country string
}

// Struct Tag
type UserForm struct {
	Name string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type name", valueType.Name())
	fmt.Println("Number fields", valueType.NumField())
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name, "with type", valueField.Type)
	}
}

func isValid(value any) (result bool) {
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface() // get data
			result = data != ""
			if result == false {
				return result // false
			}
		}
	}
	return result
}

func main() {
	readField(UserAccount{"Angga", ""})
	readField(UserAddress{"Surabaya", "Indonesia"})

	// Struct Tag
	userData := UserForm{"Angga"}
	userType := reflect.TypeOf(userData)
	userField := userType.Field(0)
	required := userField.Tag.Get("required")
	max := userField.Tag.Get("max")
	fmt.Println("required", required, "max", max)

	newUser := UserAccount{"Ari", ""}
	fmt.Println(isValid(newUser))
}