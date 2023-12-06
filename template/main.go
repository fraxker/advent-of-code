package main

import (
	"log"
	"os"
)

func main() {
	// Check args
	if len(os.Args) != 2 {
		log.Fatal("Commands: one, two")
	}
	switch os.Args[1] {
		case "one":
			One()
		case "two":
			Two()
	}
}