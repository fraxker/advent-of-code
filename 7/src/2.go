package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var HandConvertTwo = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 1,
	"T": 10,
}

func Two() {
	file, err := os.Open("/Users/asonnier/dev/other/advent-of-code-2023/7/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

	camelHands := make([]*Camel, 0)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		bid, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		hand := NewHand()
		for _, char := range strings.Split(line[0], "") {
			var num int
			if num, err = strconv.Atoi(char); err != nil {
				num = HandConvertTwo[char]
			}
			hand.Add(num)
		}

		camelHands = append(camelHands, &Camel{Bid: bid, Hand: hand})
	}


	sort.Sort(ByCamel(camelHands))

	winnings := 0
	for idx, camel := range camelHands {
		fmt.Println(camel)
		winnings += camel.Bid * (idx + 1)
	}

	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	fmt.Println(winnings)
}