package main

import (
	"log"
	"os"

	"advent.com/7/src"
)

func main() {
	// Check args
	if len(os.Args) != 2 {
		log.Fatal("Commands: one, two")
	}
	switch os.Args[1] {
		case "one":
			src.One()
		case "two":
			src.Two()
		default:
			log.Fatal("Commands: one, two")
	}
}