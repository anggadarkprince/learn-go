package main

import (
	"fmt"
	"strings"
)

func main() {
	name := " Angga  ari   "

	fmt.Println(strings.Trim(name, ""))
	fmt.Println(strings.ToLower(name))
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.Contains(name, "ari"))
	fmt.Println(strings.Split(name, " "))
	fmt.Println(strings.ReplaceAll(name, "ari", "Wijaya"))
}