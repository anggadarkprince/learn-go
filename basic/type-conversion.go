package main

import "fmt"

func main() {
	var value32 int32 = 1234567
	var value64 int64 = int64(value32)
	var value16 int16 = int16(value64) // overflow the capacity, return back (negative)
	
	fmt.Println(value32, value64, value16)

	var name = "Angga Ari Wijaya"
	var a uint8 = name[0]
	var aString = string(a)

	fmt.Println(a, aString)
}