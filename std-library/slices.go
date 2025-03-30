package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Angga", "Ari", "Wijaya", "A"}	
	values := []int{1, 2, 3}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Ari" ))
	fmt.Println(slices.Index(names, "Wijaya"))
}