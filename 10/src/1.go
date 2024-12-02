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
    pipes := make([][]rune, 0)
    for scanner.Scan() {
        line := scanner.Text()
        var pipe []rune
        for _, c := range line {
        	pipe = append(pipe, c)
        }
        pipes = append(pipes, pipe)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    // Make two dimensional array of ints same length as pipes
    disPipes := make([][]int, len(pipes))
    for i := range disPipes {
        disPipes[i] = make([]int, len(pipes[i]))   
    }

    var maxDis int
    // Find S character in pipes
    for i, pipe := range pipes {
    	for j, c := range pipe {
    		if c == 'S' {
    			maxDis, _ = TraversePipes(i, j, pipes, disPipes)
                fmt.Println(maxDis)
                break
    		}
    	}
        if maxDis != 0 {
            break
        }
    }
}