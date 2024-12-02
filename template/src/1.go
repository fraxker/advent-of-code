package src

import (
	"bufio"
	"log"
	"os"
)

func One() {
	file, err := os.Open("test1.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	scanner.Scan()
}