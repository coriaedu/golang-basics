package greetings

import "fmt"

// Hello returns a greeting for the named person
func Hello(name string) string {
	// Return a greeting that embeds the name in the message
	return fmt.Sprintf("Hello %v! Welcome!", name)
}
