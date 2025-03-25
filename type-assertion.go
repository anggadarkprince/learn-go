package main

import "fmt"

func random() interface{} {
	return "Nice"
}

func main() {
	var result any = random()
	var resultString string = result.(string)
	fmt.Println(resultString)

	//var resultInt int = result.(int)
	//fmt.Println(resultInt) // will throw panic (cannot parse to int)

	// check type with switch
	switch value := result.(type) {
	case int:
		fmt.Println("string", value) // auto converted to string
	case string:
		fmt.Println("int", value) // auto converted to int
	default:
		fmt.Println("Unknown", value)
	}
}