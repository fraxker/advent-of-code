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
    var disMap [][]int
    // Find S character in pipes
    for i, pipe := range pipes {
    	for j, c := range pipe {
    		if c == 'S' {
    			maxDis, disMap = TraversePipes(i, j, pipes, disPipes)
                break
    		}
    	}
        if maxDis != 0 {
            break
        }
    }
	

	// // Find min and max x and y of disMap
	// minX, minY, maxX, maxY := 0, 0, 0, 0
	// for i, pipe := range pipes {
	// 	for j := range pipe {
	// 		if disMap[i][j] > 0 {
	// 			if i < minX {
	// 				minX = i
	// 			}
	// 			if i > maxX {
	// 				maxX = i
	// 			}
	// 			if j < minY {
	// 				minY = j
	// 			}
	// 			if j > maxY {
	// 				maxY = j
	// 			}
	// 		}
	// 	}
	// }

	// // Prune pipes to only include those inside maxes and mins
	// pipes = pipes[minX:maxX+1]
	// for i, pipe := range pipes {
	// 	pipes[i] = pipe[minY:maxY+1]
	// }
	
	// // Print dismap
	// for _, row := range disMap {
	// 	fmt.Println(row)
	// }

	// Loop all points, check if inside and if period
	// If inside and period, add to count
	count := 0
	for i, pipe := range pipes {
		for j, c := range pipe {
			if CheckInside(disMap, i, j) && c == '.' {
				// fmt.Println(i, j)
				count++
			}
		}
	}
	fmt.Println(count)
}