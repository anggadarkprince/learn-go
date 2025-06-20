package main

import (
	"fmt"
	_ "learn-go-basic/bootstrap" // just to make init() being called without use any function
	"learn-go-basic/utilities"
)

func main() {
	result := utilities.Greeting("angga")
	//name := utilities.nameFormat("angga") // undefined nameFormat (private)
	fmt.Println(result)

	// init auto called when imported (even before Greeting())
	fmt.Println(utilities.GetConnection())
}