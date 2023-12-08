package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func One() {
	file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	scanner.Scan()

	directions := scanner.Text()

	// Skip blank line
	scanner.Scan()

	forkMap := make(map[string]*Fork)
	for scanner.Scan() {
		line := scanner.Text()
		// Parse line in this format AAA = (BBB, CCC)
		var name, left, right string
		_, err := fmt.Sscanf(line, "%3s = (%3s, %3s)\n", &name, &left, &right)
		if err != nil {
			log.Fatal(err)
		}
		forkMap[name] = NewFork(left, right)
	}

	current := "AAA"
	count := 0

	for current != "ZZZ" {
		for _, d := range directions {
			count++
			switch d {
			case 'L':
				current = forkMap[current].Left
			case 'R':
				current = forkMap[current].Right
			default:
				log.Fatal("Invalid direction")
			}

			if current == "ZZZ" {
				break
			}
		}
	}

	fmt.Println(count)
}