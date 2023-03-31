package greetings

import (
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message
	mString := fmt.Sprintf("hello, %s, it's nice to meet you", name)
	return mString
}
