package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {

	// If no name was given, return an error
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in the message
	return fmt.Sprintf(randomFormat(), name), nil
}

func randomFormat() string {
	msgFormats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return msgFormats[rand.Intn(len(msgFormats))]
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	msgs := make(map[string]string, len(names))
	var err error
	for _, name := range names {
		msgs[name], err = Hello(name)
		if err != nil {
			return nil, err
		}
	}

	return msgs, nil
}
