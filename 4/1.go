package main

import (
    "bufio"
    "strings"
    "log"
    "os"
	"fmt"
	"math"
)

func main() {
	sum := 0
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // optionally, resize scanner's capacity for lines over 64K, see next example
    for scanner.Scan() {
		count := 0
		line := strings.Split(scanner.Text(), ":")

		// 0 - Winning Numbers, 1 - Your Numbers
		card := strings.Split(line[1], "|")
		winningNumbers := strings.Fields(card[0])
		yourNumbers := strings.Fields(card[1])
		
		winningMap := make(map[string]bool)
		for _, number := range winningNumbers {
			winningMap[number] = true
		}

		for _, number := range yourNumbers {
			if _, ok := winningMap[number]; ok {
				count++
			}
		}

		if count != 0{
			sum += int(math.Pow(2, float64(count - 1)))
		}
		
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	fmt.Println(sum)
}