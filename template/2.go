package main

import (
	"bufio"
	"log"
	"os"
)

func Two() {
	file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	scanner.Scan()
}