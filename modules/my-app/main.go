package main

import (
	"fmt"

	goemailchecker "github.com/anggadarkprince/go-email-checker/v2"
)

func main() {
	// go get github.com/anggadarkprince/go-email-checker
	// go get // to update all dependencies (if we change the version in go.mod)
	isValidEmail := goemailchecker.CheckMail("angga@mail.com", false)
	fmt.Println("angga@mail.com is valid", isValidEmail)
	fmt.Println("angga@test.com is valid", goemailchecker.CheckMail("angga@test.com", true))
	fmt.Println("this-is-email is valid", goemailchecker.CheckMail("this-is-email", true))
}