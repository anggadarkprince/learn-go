package main

import "fmt"

// Nil only for some data type: interface, function, map, slice, pinter and channel
func CreateUser(name string) map[string]string { // even this declaration mention should return map, it also can be "nil"
	if name == "" {
		return nil
	}
	return map[string]string{name: name}
}

func main() {
	data := CreateUser("");
	if (data == nil) {
		fmt.Println("data is nil")
	}
	fmt.Println(data == nil)

	fmt.Println(CreateUser(""))
	fmt.Println(CreateUser("Angga"))
} 