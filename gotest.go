package main

import (
	"io"
	"log"
	"os"
)

// Just like `cat`, this program writes the contents of arg[1] to stdout
func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	check(err)
	_, err = io.Copy(os.Stdout, file)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
