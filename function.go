package main

import "fmt"

func greeting(name string) {
	fmt.Println("Hello", name)
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}

func getFullName() (string, string, string) {
	return "Angga", "Ari", "Wijaya"
}

func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "Angga"
	middleName = "Ari"
	//lastName = "Wijaya" // optional

	return firstName, middleName, lastName;
}

// Variadic (last param)
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// alias type declaration
type Filter func(string) string

func helloWithFilter(name string, filter Filter) {
	fmt.Println("Hello,", filter(name))
}

func blacklistWords(text string) string {
	if (text == "Anjing") {
		return "xxxx"
	}
	return text
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	}
	return value * factorialRecursive(value - 1)
}

func main() {
	greeting("Angga")
	greeting("Ari")

	result := adder(234, 523)
	fmt.Println(result)

	firstName, lastName, _ := getFullName()
	fmt.Println(firstName, lastName)

	fmt.Println(getCompleteName())

	fmt.Println("Total", sumAll(1, 2, 3, 4, 5))

	numbers := []int{1, 2, 3, 4, 5} // slice
	fmt.Println("Total", sumAll(numbers...)) // spread operator

	// Store to variable
	sayIt := greeting
	sayIt("Angga")

	// Function as param
	helloWithFilter("Anjing", blacklistWords)

	// Anonymous function
	helloWithFilter("Bebek", func(text string) string {
		if (text == "Bebek") {
			return "(not allowed)"
		}
		return text;
	}) // or store the function declaration into variable

	// Recursive
	fmt.Println(factorialRecursive(10))

	// Closure
	counter := 0
	increment := func() {
		fmt.Println("Increment")
		counter++ // access outside the scope
	}
	increment()
	increment()
	fmt.Println(counter)

	fmt.Println("-----------------------")
	
	// Defer: run after current function is executed
	logging := func() {
		fmt.Println("Finished (after runApp)")
	}
	runApp := func() {
		defer logging() // run after runApp finished, and keep executed even this function error
		fmt.Println("Run application")
	}
	runApp()

	fmt.Println("-----------------------")

	// Panic: stop program (defer function still will be executed)
	endApp := func() {
		fmt.Println("Finished (after run execApp)")
		message := recover()
		if (message != nil) {
			fmt.Println("Panic happen", message)
		}
	}
	execApp := func(error bool) {
		defer endApp()
		fmt.Println("Exec application")
		if (error) {
			panic("Error!!!")
		}
	}
	execApp(true)
	fmt.Println("Continue execution...") // if not "recover it will be stop"

	// Recover: get panic error message
}


