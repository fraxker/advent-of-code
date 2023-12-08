package src

import (
	"bufio"
	"fmt"
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

	directions := scanner.Text()

	// Skip blank line
	scanner.Scan()

	currents := make([]*Position, 0)
	forkMap := make(map[string]*Fork)
	for scanner.Scan() {
		line := scanner.Text()
		// Parse line in this format AAA = (BBB, CCC)
		var name, left, right string
		_, err := fmt.Sscanf(line, "%3s = (%3s, %3s)\n", &name, &left, &right)
		if err != nil {
			log.Fatal(err)
		}
		if name[2] == 'A' {
			currents = append(currents, &Position{Start: name, Position: name, AtEnd: false})
		}
		forkMap[name] = NewFork(left, right)
	}

	count := 0

	cycles := make([]int, 0)

	for len(cycles) != len(currents) {
		for _, d := range directions {
			count++
			if count % 100_000_000 == 0 {
				fmt.Println(count)
			}
			for _, c := range currents {
				switch d {
				case 'L':
					c.Position = forkMap[c.Position].Left
				case 'R':
					c.Position = forkMap[c.Position].Right
				default:
					log.Fatal("Invalid direction")
				}
				if !c.AtEnd && c.Position[2] == 'Z' {
					cycles = append(cycles, count)
					fmt.Println(c.Start, count)
					c.AtEnd = true
				}
			}
		}
	}

	fmt.Println(LCM(cycles[0], cycles[1], cycles[2:]...))

}