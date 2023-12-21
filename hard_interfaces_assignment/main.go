package main

import (
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("Filename was not provided")
	}

	fn := args[1]
	fp, err := os.Open(fn)

	if err != nil {
		log.Fatal("File ", fn, " could not be opened. Error: ", err)
	}

	io.Copy(os.Stdout, fp)
}
