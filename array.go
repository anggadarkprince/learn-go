package main

import "fmt"

func main() {
	// array declaration
	var names [3]string
	names[0] = "Angga"
	names[1] = "Ari"
	names[2] = "Wijaya"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	names[2] = "Kusuma";
	//names[3] = "Keenan"; error

	// direct assignment
	var scores = [5]int{
		70, 80, 90, // index 3 and 4 will be 0 (depend on data type, empty string for string)
	}
	fmt.Println(scores)

	fmt.Println(len(scores))

	// dynamic length based on direct assignment values
	var values = [...]int{
		1, 2, 3, 4, 5,
	}
	fmt.Println(values)

}