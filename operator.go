package main

import "fmt"

func main() {
	// Math operator
	var x = 10
	var y = 3
	var total = x + y
	fmt.Println(total)

	var sub = x - y
	fmt.Println(sub)

	var ex = x * y
	fmt.Println(ex)

	var diff = x / y
	fmt.Println(diff)
	fmt.Println(ex)

	var mod = x % y
	fmt.Println(mod)

	// Augmented operator
	var a = 10
	a += 5
	fmt.Println(a)
	
	var b = 10
	b *= 5
	fmt.Println(b)

	// Unary operator
	var c = 5
	c++
	fmt.Println(c)
	c--
	fmt.Println(c)

	var isTrue = false
	fmt.Println(!isTrue)

	// Comparison operator (<, >, <=, >=, ==, !=)
	var isLarger = 2 > 4
	var isDifferent = 2 != 3
	var isSame = "angga" == "ari"
	fmt.Println(isLarger, isDifferent, isSame)

	// Boolean operator (&&, ||, !)
	const userId = 1
	const permission = "view-order"
	var hasPemission = userId == 1 && permission == "view-user"
	fmt.Println(hasPemission)
	fmt.Println(!hasPemission)
}