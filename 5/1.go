package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"advent.com/5/util"
)

func One() {
	lowest := int64(-1)
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

	// Scan in seed numbers
	scanner.Scan()
	seed_line := strings.Split(scanner.Text(), ":")
	seed_numbers := strings.Fields(seed_line[1])

	// Skip blank line
	scanner.Scan()
	maps := make([]*util.RangeMap, 7)
	for i := 0; i < len(maps); i++ {
		// Skip header
		scanner.Scan()
		maps[i] = util.NewRangeMap()
		// scan in entries, break on blank line
		for scanner.Scan() {
			temp := scanner.Text()
			line := strings.Fields(temp)
			if len(line) == 0 {
				break
			}
			source, err := strconv.ParseInt(line[0], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			destination, err := strconv.ParseInt(line[1], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			width, err := strconv.ParseInt(line[2], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			maps[i].AddEntry(util.NewRangeEntry(source, destination, width))
		}
	}

	// Convert Seed numbers through map
	for _, seed := range seed_numbers {
		seed_int, err := strconv.ParseInt(seed, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		for _, m := range maps {
			fmt.Printf("%d -> ", seed_int)
			seed_int = m.MapSource(seed_int)
			fmt.Printf("%d\n", seed_int)
		}
		if lowest == -1 || seed_int < lowest {
			lowest = seed_int
		}
	}

	fmt.Println(lowest)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}