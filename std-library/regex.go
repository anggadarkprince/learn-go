package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`a([a-z])i`)
	fmt.Println(regex.MatchString("ari"));
	fmt.Println(regex.MatchString("aki"));
	fmt.Println(regex.MatchString("a3i"));

	fmt.Println(regex.FindAllString("angga ari ani aHi", 10)) // take match max 10
}