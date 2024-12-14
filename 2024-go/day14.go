package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day14() {
	// sample := []string{
	// 	"p=0,4 v=3,-3",
	// 	"p=6,3 v=-1,-3",
	// 	"p=10,3 v=-1,2",
	// 	"p=2,0 v=2,-1",
	// 	"p=0,0 v=1,3",
	// 	"p=3,0 v=-2,-2",
	// 	"p=7,6 v=-1,-3",
	// 	"p=3,0 v=-1,-2",
	// 	"p=9,3 v=2,3",
	// 	"p=7,3 v=-1,2",
	// 	"p=2,4 v=2,-3",
	// 	"p=9,5 v=-3,-3",
	// }
	input := ReadLines("input/day14.txt")
	width, height := 101, 103
	robots := day14_parse_robots(input)
	seconds := 100

	for i, robot := range robots {
		robots[i].pos.col = Mod((robot.pos.col + robot.vel.col*seconds), width)
		robots[i].pos.row = Mod((robot.pos.row + robot.vel.row*seconds), height)
	}
	safetyFactor := day14_safety_factor(robots, width, height)

	fmt.Println("Day 14 Part 01:", safetyFactor)

	robots = day14_parse_robots(input)
	i := 0
	for {
		longestLine := day14_longest_line(robots, width, height)
		if longestLine > width/4 {
			// day14_print(robots, width, height)
			break
		}
		i++
		for j, robot := range robots {
			robots[j].pos.col = Mod((robot.pos.col + robot.vel.col), width)
			robots[j].pos.row = Mod((robot.pos.row + robot.vel.row), height)
		}
	}

	fmt.Println("Day 14 Part 02:", i)
}

func day14_longest_line(robots []Robot, width, height int) int {
	grid := make([][]bool, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]bool, width)
	}

	for _, robot := range robots {
		grid[robot.pos.row][robot.pos.col] = true
	}

	maxLine := 0

	// Check horizontal lines
	for row := 0; row < height; row++ {
		currentLine := 0
		for col := 0; col < width; col++ {
			if grid[row][col] {
				currentLine++
				if currentLine > maxLine {
					maxLine = currentLine
				}
			} else {
				currentLine = 0
			}
		}
	}

	// Check vertical lines
	for col := 0; col < width; col++ {
		currentLine := 0
		for row := 0; row < height; row++ {
			if grid[row][col] {
				currentLine++
				if currentLine > maxLine {
					maxLine = currentLine
				}
			} else {
				currentLine = 0
			}
		}
	}

	return maxLine
}

// func day14_print(robots []Robot, width, height int) {
// 	grid := make([][]string, height)
// 	for i := 0; i < height; i++ {
// 		grid[i] = make([]string, width)
// 		for j := 0; j < width; j++ {
// 			grid[i][j] = "."
// 		}
// 	}
// 	for _, robot := range robots {
// 		grid[robot.pos.row][robot.pos.col] = "#"
// 	}
// 	for _, row := range grid {
// 		fmt.Println(strings.Join(row, ""))
// 	}
// }

func day14_safety_factor(robots []Robot, width, height int) int {
	quadrants := [4]int{}
	for _, robot := range robots {
		quadrant := day14_quadrant(robot, width, height)
		if quadrant != -1 {
			quadrants[quadrant]++
		}
	}
	safetyFactor := 1
	for _, quadrant := range quadrants {
		safetyFactor *= quadrant
	}
	return safetyFactor
}

func day14_quadrant(robot Robot, width, height int) int {
	halfWidth := width / 2
	halfHeight := height / 2
	if robot.pos.col < halfWidth {
		if robot.pos.row < halfHeight {
			return 0
		} else if robot.pos.row > halfHeight {
			return 3
		}
	} else if robot.pos.col > halfWidth {
		if robot.pos.row < halfHeight {
			return 1
		} else if robot.pos.row > halfHeight {
			return 2
		}
	}
	return -1
}

type Robot struct {
	pos Point
	vel Point
}

func day14_parse_robots(input []string) []Robot {
	robots := []Robot{}
	for _, line := range input {
		robot := Robot{}
		parts := strings.Split(line, " ")
		posParts := strings.Split(parts[0][2:], ",")
		velParts := strings.Split(parts[1][2:], ",")
		robot.pos.col, _ = strconv.Atoi(posParts[0])
		robot.pos.row, _ = strconv.Atoi(posParts[1])
		robot.vel.col, _ = strconv.Atoi(velParts[0])
		robot.vel.row, _ = strconv.Atoi(velParts[1])
		robots = append(robots, robot)
	}
	return robots
}
