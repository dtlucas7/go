package main

import (
	"fmt"
	"log"
	"os"
)

func ls(path string) {

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	// iterate over the files and print their names
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func main() {
	ls(".")
}
