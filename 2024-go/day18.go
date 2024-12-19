package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day18() {
	// sample := []string{
	// 	"5,4",
	// 	"4,2",
	// 	"4,5",
	// 	"3,0",
	// 	"2,1",
	// 	"6,3",
	// 	"2,4",
	// 	"1,5",
	// 	"0,6",
	// 	"3,3",
	// 	"2,6",
	// 	"5,1",
	// 	"1,2",
	// 	"5,5",
	// 	"2,5",
	// 	"6,5",
	// 	"1,4",
	// 	"0,4",
	// 	"6,4",
	// 	"1,1",
	// 	"6,1",
	// 	"1,0",
	// 	"0,5",
	// 	"1,6",
	// 	"2,0",
	// }
	input := ReadLines("input/day18.txt")
	points := day18_parse(input)
	width, height, fallen := 71, 71, 1024
	grid := day18_grid(width, height)
	for _, point := range points[:fallen] {
		grid[point.row][point.col] = '#'
	}
	start, end := Point{0, 0}, Point{height - 1, width - 1}
	steps := day18_minimum_steps(grid, start, end)

	fmt.Println("Day 18 Part 01:", steps)

	var blocker Point
	left, right := 0, len(points)-1
	for left <= right {
		mid := (left + right) / 2

		grid = day18_grid(width, height)
		for _, point := range points[:mid+1] {
			grid[point.row][point.col] = '#'
		}

		steps = day18_minimum_steps(grid, start, end)
		if steps == -1 {
			blocker = points[mid]
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	result := strconv.Itoa(blocker.col) + "," + strconv.Itoa(blocker.row)

	fmt.Println("Day 18 Part 02:", result)
}

func day18_minimum_steps(grid [][]byte, start, end Point) int {
	queue := []Point{start}
	visited := map[Point]int{start: 0}

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		steps := visited[point]

		if point == end {
			return steps
		}

		for _, next := range Neighbors(point) {
			if next.row < 0 || next.row >= len(grid) || next.col < 0 || next.col >= len(grid[0]) {
				continue
			}
			if grid[next.row][next.col] == '#' {
				continue
			}
			if _, ok := visited[next]; !ok {
				visited[next] = steps + 1
				queue = append(queue, next)
			}
		}
	}
	return -1
}

func day18_grid(width, height int) [][]byte {
	grid := make([][]byte, height)
	for i := range grid {
		grid[i] = make([]byte, width)
	}
	return grid
}

func day18_parse(input []string) []Point {
	points := []Point{}
	for _, line := range input {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, Point{y, x})
	}
	return points
}
