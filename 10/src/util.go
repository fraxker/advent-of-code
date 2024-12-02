package src

type Pipe struct {
	x, y, prevX, PrevY int
	pipe rune
}

// Function to return the max distance of loop from start
func TraversePipes(startX, startY int, pipes [][]rune, disPipes [][]int) (int, [][]int) {
	pipesToCheck := make([]*Pipe, 0)
	pipesToCheck = append(pipesToCheck, &Pipe{x: startX, y: startY, prevX: startX, PrevY: startY, pipe: pipes[startX][startY]})
	disPipes[startX][startY] = 0
	for len(pipesToCheck) > 0 {
		var pipe *Pipe
		pipe, pipesToCheck = pipesToCheck[0], pipesToCheck[1:]
		if pipe.pipe == 'S' {
			// Check if prex and prey don't equal x and y
			if pipe.prevX != pipe.x || pipe.PrevY != pipe.y {
				disPipes[pipe.x][pipe.y] = disPipes[pipe.prevX][pipe.PrevY] + 1
				return (disPipes[pipe.prevX][pipe.PrevY] + 1) / 2, disPipes
			}
			// Check all 4 directions
			// Check North for |, F, or 7
			if pipe.x > 0 {
				upPipe := pipes[pipe.x-1][pipe.y]
				if upPipe == '|' || upPipe == 'F' || upPipe == '7' {
					disPipes[pipe.x-1][pipe.y] = disPipes[pipe.x][pipe.y] + 1
					pipesToCheck = append(pipesToCheck, &Pipe{x: pipe.x - 1, y: pipe.y, prevX: pipe.x, PrevY: pipe.y, pipe: upPipe})
					continue
				}
			}

			// Check South for |, J, or L
			if pipe.x < len(pipes)-1 {
				downPipe := pipes[pipe.x+1][pipe.y]
				if downPipe == '|' || downPipe == 'J' || downPipe == 'L' {
					disPipes[pipe.x+1][pipe.y] = disPipes[pipe.x][pipe.y] + 1
					pipesToCheck = append(pipesToCheck, &Pipe{x: pipe.x + 1, y: pipe.y, prevX: pipe.x, PrevY: pipe.y, pipe: downPipe})
					continue
				}
			}

			// Check East for -, J, or 7
			if pipe.y < len(pipes[0])-1 {
				rightPipe := pipes[pipe.x][pipe.y+1]
				if rightPipe == '-' || rightPipe == 'J' || rightPipe == '7' {
					disPipes[pipe.x][pipe.y+1] = disPipes[pipe.x][pipe.y] + 1
					pipesToCheck = append(pipesToCheck, &Pipe{x: pipe.x, y: pipe.y + 1, prevX: pipe.x, PrevY: pipe.y, pipe: rightPipe})
					continue
				}
			}

			// Check West for -, F, or L
			if pipe.y > 0 {
				leftPipe := pipes[pipe.x][pipe.y-1]
				if leftPipe == '-' || leftPipe == 'F' || leftPipe == 'L' {
					disPipes[pipe.x][pipe.y-1] = disPipes[pipe.x][pipe.y] + 1
					pipesToCheck = append(pipesToCheck, &Pipe{x: pipe.x, y: pipe.y - 1, prevX: pipe.x, PrevY: pipe.y, pipe: leftPipe})
					continue
				}
			}
		}

		// | has connections above and below
		if pipe.pipe == '|' {
			// Check is previous pipe was below
			if pipe.prevX > pipe.x {
				// Add above to pipesToCheck
				upPipe := pipes[pipe.x-1][pipe.y]
				disPipes[pipe.x-1][pipe.y] = disPipes[pipe.x][pipe.y] + 1
				pipesToCheck = append(pipesToCheck, &Pipe{x: pipe.x - 1, y: pipe.y, prevX: pipe.x, PrevY: pipe.y, pipe: upPipe})
			} else {
				// Add below to pipesToCheck
				downPipe := pipes[pipe.x+1][pipe.y]
				disPipes[pipe.x+1][pipe.y] = disPipes[pipe.x][pipe.y] + 1
				pipesToCheck = append(pipesToCheck, &Pipe{x: pipe.x + 1, y: pipe.y, prevX: pipe.x, PrevY: pipe.y, pipe: downPipe})
			}
			continue
		}

		// - has connections to the left and right
		if pipe.pipe == '-' {
			// Check is previous pipe was to the right
			if pipe.PrevY > pipe.y {
				// Add left to pipesToCheck
				leftPipe := pipes[pipe.x][pipe.y-1]
				disPipes[pipe.x][pipe.y-1] = disPipes[pipe.x][pipe.y] + 1
				pipesToCheck = append(pipesToCheck, &Pipe{x: pipe.x, y: pipe.y - 1, prevX: pipe.x, PrevY: pipe.y, pipe: leftPipe})
			} else {
				// Add right to pipesToCheck
				rightPipe := pipes[pipe.x][pipe.y+1]
				disPipes[pipe.x][pipe.y+1] = disPipes[pipe.x][pipe.y] + 1
				pipesToCheck = append(pipesToCheck, &Pipe{x: pipe.x, y: pipe.y + 1, prevX: pipe.x, PrevY: pipe.y, pipe: rightPipe})
			}
			continue
		}

		// F has connections below and right
		if pipe.pipe == 'F' {
			// If previous pipe was to the below, add right
			if pipe.prevX > pipe.x {
				rightPipe := pipes[pipe.x][pipe.y+1]
				disPipes[pipe.x][pipe.y+1] = disPipes[pipe.x][pipe.y] + 1
				pipesToCheck = append(pipesToCheck, &Pipe{x: pipe.x, y: pipe.y + 1, prevX: pipe.x, PrevY: pipe.y, pipe: rightPipe})
			} else {
				// Add below
				downPipe := pipes[pipe.x+1][pipe.y]
				disPipes[pipe.x+1][pipe.y] = disPipes[pipe.x][pipe.y] + 1
				pipesToCheck = append(pipesToCheck, &Pipe{x: pipe.x + 1, y: pipe.y, prevX: pipe.x, PrevY: pipe.y, pipe: downPipe})
			}
			continue
		}
		
		// J has connections above and left
		if pipe.pipe == 'J' {
			// If previous pipe was above, add left
			if pipe.prevX < pipe.x {
				leftPipe := pipes[pipe.x][pipe.y-1]
				disPipes[pipe.x][pipe.y-1] = disPipes[pipe.x][pipe.y] + 1
				pipesToCheck = append(pipesToCheck, &Pipe{x: pipe.x, y: pipe.y - 1, prevX: pipe.x, PrevY: pipe.y, pipe: leftPipe})
			} else {
				// Add above
				upPipe := pipes[pipe.x-1][pipe.y]
				disPipes[pipe.x-1][pipe.y] = disPipes[pipe.x][pipe.y] + 1
				pipesToCheck = append(pipesToCheck, &Pipe{x: pipe.x - 1, y: pipe.y, prevX: pipe.x, PrevY: pipe.y, pipe: upPipe})
			}
			continue
		}

		// L has connections above and right
		if pipe.pipe == 'L' {
			// If previous pipe was above, add right
			if pipe.prevX < pipe.x {
				rightPipe := pipes[pipe.x][pipe.y+1]
				disPipes[pipe.x][pipe.y+1] = disPipes[pipe.x][pipe.y] + 1
				pipesToCheck = append(pipesToCheck, &Pipe{x: pipe.x, y: pipe.y + 1, prevX: pipe.x, PrevY: pipe.y, pipe: rightPipe})
			} else {
				// Add above
				upPipe := pipes[pipe.x-1][pipe.y]
				disPipes[pipe.x-1][pipe.y] = disPipes[pipe.x][pipe.y] + 1
				pipesToCheck = append(pipesToCheck, &Pipe{x: pipe.x - 1, y: pipe.y, prevX: pipe.x, PrevY: pipe.y, pipe: upPipe})
			}
			continue
		}

		// 7 has connections below and left
		if pipe.pipe == '7' {
			// If previous pipe was below, add left
			if pipe.prevX > pipe.x {
				leftPipe := pipes[pipe.x][pipe.y-1]
				disPipes[pipe.x][pipe.y-1] = disPipes[pipe.x][pipe.y] + 1
				pipesToCheck = append(pipesToCheck, &Pipe{x: pipe.x, y: pipe.y - 1, prevX: pipe.x, PrevY: pipe.y, pipe: leftPipe})
			} else {
				// Add below
				downPipe := pipes[pipe.x+1][pipe.y]
				disPipes[pipe.x+1][pipe.y] = disPipes[pipe.x][pipe.y] + 1
				pipesToCheck = append(pipesToCheck, &Pipe{x: pipe.x + 1, y: pipe.y, prevX: pipe.x, PrevY: pipe.y, pipe: downPipe})
			}
			continue
		}
	}
	return 0, nil
}

func CheckInside(disMap [][]int, x, y int) bool {
	// Check in ray in all 4 directions starting from x, y
	// If all four have a non zero value, then x, y is inside
	// If any of the four have a zero value, then x, y is outside
	
	if disMap[x][y] != 0 {
		return false
	}
	// Check North
	boundNorth := 0
	for i := x; i >= 0; i-- {
		if disMap[i][y] != 0 {
			boundNorth++
		}
	}
	if boundNorth == 0 || boundNorth % 2 == 0 {
		return false
	}

	// Check South
	boundSouth := 0
	for i := x; i < len(disMap); i++ {
		if disMap[i][y] != 0 {
			boundSouth++
		}
	}
	if boundSouth == 0 || boundSouth % 2 == 0 {
		return false
	}

	// Check East
	boundEast := 0
	for i := y; i < len(disMap[0]); i++ {
		if disMap[x][i] != 0 {
			boundEast++
		}
	}
	if boundEast == 0 || boundEast % 2 == 0 {
		return false
	}

	// Check West
	boundWest := 0
	for i := y; i >= 0; i-- {
		if disMap[x][i] != 0 {
			boundWest++
		}
	}
	if boundWest == 0 || boundWest % 2 == 0 {
		return false
	}

	return true
}