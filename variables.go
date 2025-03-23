package main

import "fmt"

func main() {
	var name string

	name = "Angga Ari Wijaya"
	fmt.Println(name)
	
	name = "Keenan Evander"
	fmt.Println(name)

	var age = 30 // without mention the data type
	fmt.Println("Age", age);

	location := "Surabaya, Indonesia" // shortcut for var initialization
	fmt.Println("Location", location)

	var (
		fullName = "Angga Ari W."
		city = "Surabaya"
		country = "Indonesia"
	)

	fmt.Println(fullName, city, country)
}