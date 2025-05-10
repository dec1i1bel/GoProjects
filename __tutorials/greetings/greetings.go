package greetings

import (
	"fmt"
)

// Returns a greeting for the named person
func Hello(name string) string {
	if name == "" {
		return ""
	}

	message := fmt.Sprintf("Hi, %v. welcome!", name)
	return message
}
