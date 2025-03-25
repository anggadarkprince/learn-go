package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// by default Go using "Pass by value" not "Pass by reference", so when we assign variable from another variable it's copy by default

func main() {
	address1 := Address{"Surabaya", "East Java", "Indonesia"}
	address2 := address1 // copy the value (clone)

	address2.City = "Gresik" // not change address1

	fmt.Println(address1)
	fmt.Println(address2)

	// what if I need by reference, then we can use Pointer
	// we use operator & following with the variable name
	address3 := &address1 // address3 is a pointer to address1, so when we modify address3 then address1 will be change

	address3.City = "Sidoarjo"
	fmt.Println(address3)
	fmt.Println(address1) // also changed

	// when we use & operator, it only change what we reference
	// if we want to change any variable operator *
	//address3 = Address{"Malang", "East Java", "Indonesia"} // cannot re-assign a pointer
	address3 = &Address{"Malang", "East Java", "Indonesia"}
	fmt.Println(address3)
	fmt.Println(address1)

}