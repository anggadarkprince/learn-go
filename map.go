package main

import "fmt"

func main() {
	user := map[string]string{
		"name": "Angga",
		"username": "angga.ari",
		"status": "active",
	}
	fmt.Println(user)
	fmt.Println(user["name"])
	fmt.Println(user["username"])
	fmt.Println(user["age"]) // if not exist return default value (depend on data type of value)

	// Create new empty map
	book := make(map[string]string)
	book["title"] = "Pragmatic Programmer"
	fmt.Println(book)

	// Map size
	fmt.Println(len(user))

	// Add or Update value
	user["status"] = "inactive"
	fmt.Println(user)

	// Delete key
	delete(user, "status")
	fmt.Println(user)
}