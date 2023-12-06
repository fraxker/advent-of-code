package src

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line1 := strings.Split(scanner.Text(), ":")
	race_length := strings.Join(strings.Fields(line1[1]), "")
	scanner.Scan()
	line2 := strings.Split(scanner.Text(), ":")
	race_record := strings.Join(strings.Fields(line2[1]), "")

	length, err := strconv.Atoi(race_length)
	if err != nil {
		log.Fatal(err)
	}
	record, err := strconv.Atoi(race_record)
	if err != nil {
		log.Fatal(err)
	}

	intersections := Solve_Quadratic(-1, float64(length), -float64(record))
	var range_int int
	if len(intersections) == 2 {
		range_int = int(math.Ceil(intersections[1])) - int(math.Floor(intersections[0])) - 1
	} else if len(intersections) == 1 {
		range_int = int(math.Ceil(intersections[0])) - int(math.Floor(intersections[0])) - 1
	}
	fmt.Printf("length: %d, record: %d, intecections: %v, range: %d\n", length, record, intersections, range_int)
}