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

func main() {
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
	expanded_seed_numbers := []*util.SeedPair{}
	// Pair seeds into groups of two, 1st is start point, 2nd is length, add all numbers to expanded_see_numbers
	for i := 0; i < len(seed_numbers); i += 2 {
		start, err := strconv.ParseInt(seed_numbers[i], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		length, err := strconv.ParseInt(seed_numbers[i + 1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		expanded_seed_numbers = append(expanded_seed_numbers, &util.SeedPair{Start: start, Length: length})
	}

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

	cache := make(map[int64]int64)

	// Convert Seed numbers through map
	for _, seed := range expanded_seed_numbers {
		fmt.Printf("Seed: %d, %d\n", seed.Start, seed.Length)
		for i := seed.Start - 1; i <= seed.Start + seed.Length; i++ {
			seed_num := i
			if _, ok := cache[seed_num]; !ok {
				for _, m := range maps {
					seed_num = m.MapSource(seed_num)
				}
				if lowest == -1 || seed_num < lowest {
					fmt.Println(i, seed_num)
					lowest = seed_num
				}
			} 
		}
	}

	fmt.Println(lowest)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}