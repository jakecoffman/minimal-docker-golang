package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Just like `cat`, this program writes the contents of arg[1] to stdout
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: minimal-docker-golang file-to-cat")
		return
	}
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
