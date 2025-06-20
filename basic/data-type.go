package main

import "fmt"

func main() {
	// String
	fmt.Println("Hello");
	fmt.Println("Total 'hello' chars", len("Hello"));
	fmt.Println("Get first char of 'hello'", "Hello"[0]); // byte format

	// Number
	fmt.Println("Satu = ", 1);
	fmt.Println("Dua = ", 2);
	fmt.Println("Tiga koma lima = ", 3.5);

	// Boolean
	fmt.Println("Benar = ", true);
	fmt.Println("Salah = ", false);
}