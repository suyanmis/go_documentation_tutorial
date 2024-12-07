package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

// a function whose name starts with a capital letter can be called by a function not in the same package.
// exported name
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := randomFormat(name)
	return message, nil
}

func randomFormat(name string) string {

	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Welcome back, %v",
		"Ohayou, %v",
	}
	randomValue := rand.Intn(len(formats))

	return fmt.Sprintf(formats[randomValue], name)
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	if len(names) == 0 {
		return nil, errors.New("names slice is empty")
	}
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
