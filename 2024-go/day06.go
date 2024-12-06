package main

import "fmt"

func day06() {
	// sample := []string{
	// 	"....#.....",
	// 	".........#",
	// 	"..........",
	// 	"..#.......",
	// 	".......#..",
	// 	"..........",
	// 	".#..^.....",
	// 	"........#.",
	// 	"#.........",
	// 	"......#...",
	// }
	input := ReadLines("input/day06.txt")
	start := FindInGrid(input, '^')
	dir := Point{-1, 0}
	path := day06_get_path(input, start, dir)

	fmt.Println("Day 06 Part 01:", len(path))

	count := 0
	for obstacle := range path {
		if day06_is_loop(input, start, obstacle, dir) {
			count++
		}
	}

	fmt.Println("Day 06 Part 02:", count)
}

func day06_is_loop(input []string, point, obstacle Point, dir Point) bool {
	visited := map[Point]int{point: 0}
	for {
		next := point.Add(dir)
		if !InsideGrid(input, next) {
			break
		}
		if input[next.row][next.col] != '#' && next != obstacle {
			if visited[next] > 4 {
				return true
			}
			point = next
			visited[point]++
		} else {
			dir.TurnRight()
		}
	}
	return false
}

func day06_get_path(input []string, point Point, dir Point) map[Point]int {
	visited := map[Point]int{point: 0}
	for {
		next := point.Add(dir)
		if !InsideGrid(input, next) {
			break
		}
		if input[next.row][next.col] != '#' {
			point = next
			visited[point]++
		} else {
			dir.TurnRight()
		}
	}
	return visited
}
