package main

import "fmt"

func main() {
	type CustomerNumber string; // alias of string

	var customer1 CustomerNumber = "1123"

	var nextCustomer = "5678"
	var custNo = CustomerNumber(nextCustomer)
	
	fmt.Println(customer1, custNo);
}