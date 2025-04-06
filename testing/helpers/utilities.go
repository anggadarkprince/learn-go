package helpers

import (
	"strings"
)

func GreetingTo(name string) string {
	if name == "" {
		return "Hello there"
	}
	return "Hello " + name
}

func StringToSlug(values ...any) string {
	var parts []string
	for _, val := range values {
		switch v := val.(type) {
		case string:
			parts = append(parts, strings.ToLower(strings.ReplaceAll(v, " ", "-")))
		case []string:
			for _, s := range v {
				parts = append(parts, strings.ToLower(strings.ReplaceAll(s, " ", "-")))
			}
		}
	}

	return strings.Join(parts, "-")
}

func StringContains(value string, search string) bool {
	return strings.Contains(strings.ToLower(value), strings.ToLower(search));
}