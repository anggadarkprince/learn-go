package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func (address *Address) Clear() { // add * will automatic make golang passing by reference, it's recommended for struct method
	address.City = ""
	address.Province = ""
	address.Country = ""
}

func ChangeAddressToIndonesia(address *Address) { // should pointer to be passed
	address.Country = "Indonesia"
}

// by default Go using "Pass by value" not "Pass by reference", so when we assign variable from another variable it's copy by default

func main() {
	address1 := Address{"Surabaya", "East Java", "Indonesia"}
	address2 := address1 // copy the value (clone)

	address2.City = "Gresik" // not change address1

	fmt.Println(address1)
	fmt.Println(address2)

	fmt.Println("--------- & operator pointer")

	// what if I need by reference, then we can use Pointer
	// we use operator & following with the variable name
	var address3 *Address = &address1 // address3 is a pointer to address1, so when we modify address3 then address1 will be change
	addressX := address3 // not copy since the address3 it's pointer
	addressX.City = "Surabaya"
	fmt.Println(addressX, address3)

	address3.City = "Sidoarjo"
	fmt.Println(address3)
	fmt.Println(address1) // also changed

	fmt.Println("--------- & operator new value")

	// when we use & operator, it only change what we reference
	//address3 = Address{"Malang", "East Java", "Indonesia"} // cannot re-assign a pointer
	address3 = &Address{"Malang", "East Java", "Indonesia"} // address 3 now have new reference but with new value (move pointer to new value)
	fmt.Println(address3)
	fmt.Println(address1) // original address won't change, because we move pointer address3 from address1 to new value 

	fmt.Println("--------- * operator")

	// Dereferencing a Pointer: change the value from pointer
	// * operator used to get the value stored at the pointer’s address.
	address4 := &address1
	*address4 = Address{"Gresik", "East Java", "Indonesia"} // Changing value at the address stored in address4
	fmt.Println(address4)
	fmt.Println(address1) // original also changed 

	fmt.Println("--------- new operator")
	newAddress := new(Address) // empty pointer variable
	refAddress := newAddress // no need & because newAddress it's already a pointer
	refAddress.Country = "Singapore"
	fmt.Println(newAddress, refAddress)

	fmt.Println("--------- pointer in function")
	//var foreignAddress *Address = &Address{"Sidney", "State", "Australia"}
	foreignAddress := Address{"Sidney", "State", "Australia"}
	ChangeAddressToIndonesia(&foreignAddress) //change variable to pointer
	fmt.Println(foreignAddress)

	fmt.Println("--------- pointer in struct method")
	foreignAddress.Clear()
	fmt.Println("After Clear()", foreignAddress)
}