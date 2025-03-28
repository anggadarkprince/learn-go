package utilities

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Greeting(name string) string {
	return "Hi, " + nameFormat(name)
}

// lowercase in front function name is non-exportable (private)
func nameFormat(name string) string {
	return cases.Title(language.Indonesian).String(name);
}