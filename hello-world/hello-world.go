package main

import (
	"fmt"

	"github.com/dtlucas/go/greetings/greetings"
)

func main() {
	// note the syntax of calling the `quote` package
	fmt.Println(greetings.Hello())
	fmt.Println()
}
