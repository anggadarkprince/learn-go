package main

import "fmt"

func main() {
	// Basic loop with condition
	counter := 1
	for counter <= 10 {
		fmt.Println("Iteration", counter)
		if counter == 5 {
			fmt.Println("Break!")
			break;
		}
		counter++
	}

	counter = 1
	for {
		if counter > 3 {
			break
		}
		fmt.Println("Alt iteration", counter)
		counter++
	}

	// For with statement (init and post)
	for counter := 1; counter <= 5; counter++ {
		if counter % 2 == 0 {
			continue; // skip code and continue the loop
		}
		fmt.Println("Loop", counter)
	}

	// Loop data collection: array, slice, map
	names := []string{"Angga", "Ari", "Wijaya"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	user := map[string]string{
		"name": "Angga Ari Wijaya",
		"email": "angga@mail.com",
		"dob": "1992-05-26",
	}
	for key, value := range user {
		fmt.Println(key, value)
	}

}