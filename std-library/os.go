package main

import (
	"fmt"
	"os"
)

func main() {
	// Args
	args := os.Args
	fmt.Println(args) // go run os.go Angga "Ari wijaya" 20
	for _, arg := range args {
		fmt.Println(arg)
	}

	// Hostname
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("error", err.Error())
	}

}