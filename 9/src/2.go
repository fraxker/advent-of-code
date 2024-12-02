package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Two() {
	file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	sum := 0
    scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		
		numbers := make([]int, len(fields))
		for i, field := range fields {
			numbers[i], err = strconv.Atoi(field)
			if err != nil {
				log.Fatal(err)
			}
		}
		sum += PredictPrev(numbers)
	}

	fmt.Println(sum)
}