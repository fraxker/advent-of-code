package main

import (
	// "log"
	"os"

	"advent.com/10/src"
)

func main() {
	// Check args
	// if len(os.Args) != 2 {
	// 	log.Fatal("Commands: one, two")
	// }
	// src.Two()
	switch os.Args[1] {
		case "one":
			src.One()
		case "two":
			src.Two()
	}
}