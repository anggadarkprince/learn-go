package main

import (
	"fmt"
	"strconv"
)

func main() {
	booleanVal, err := strconv.ParseBool("1")
	if err == nil {
		fmt.Println("boolean", booleanVal)
	} else {
		fmt.Println("Error", err.Error())
	}
	
	intVal, err := strconv.Atoi("1000")
	if err == nil {
		fmt.Println("integer", intVal)
	} else {
		fmt.Println("Error", err.Error())
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	var stringInt = strconv.Itoa(999)
	fmt.Println(stringInt)
}