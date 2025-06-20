package main

import "fmt"

func main() {
	const firstName = "Angga"
	const gender = "male"
	const location = "Surabaya, Indonesia" // allow unused

	const (
		city = "Surabaya"
		country = "Indonesia"
	)

	// gender = "female" // error, cannot re-assign value

	fmt.Println(firstName, gender)
}