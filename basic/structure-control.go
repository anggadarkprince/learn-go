package main

import "fmt"

func main() {
	// IF
	score := 70
	if (score >= 80) {
		fmt.Println("Passed")
	} else if(score >= 60) {
		fmt.Println("Remedy")
	} else {
		fmt.Println("Not Passed")
	}

	// If without braces
	if score == 100 {
		fmt.Println("Perfect")
	} else {
		fmt.Println("So so")
	}

	// Assign and check
	name := "Angga Ari Wijaya"
	if length := len(name); length > 5 {
		fmt.Println("Name is too long, yours", length, "characters")
	}


	// SWITCH
	switch score {
	case 100:
		fmt.Println("Perfect score") // no break; statement
	case 80:
		fmt.Println("Lucky number")
	default:
		fmt.Println("Your score", score)
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name is too long, yours", length, "characters")
	case false:
		fmt.Println("Name is valid ")
	}

	switch {
	case score > 90:
		fmt.Println("Perfect")
	case score >= 80:
		fmt.Println("Passed")
	default:
		fmt.Println("Not passed")
	}
}