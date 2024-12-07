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
	length, path := day06_get_path(input, start, dir)

	fmt.Println("Day 06 Part 01:", length)

	count := 0
	visited := map[Point]bool{}
	for i, obstacle := range path {
		if _, v := visited[obstacle.point]; !v {
			visited[obstacle.point] = true
			if i > 0 && day06_is_loop(input, path[i-1].point, obstacle.point, obstacle.dir) {
				count++
			}
		}
	}

	fmt.Println("Day 06 Part 02:", count)
}

type visitor struct {
	point Point
	dir   Point
}

func day06_is_loop(input []string, point, obstacle Point, dir Point) bool {
	visited := map[visitor]bool{{point, dir}: false}
	for {
		next := point.Add(dir)
		if !InsideGrid(input, next) {
			break
		}
		if _, v := visited[visitor{next, dir}]; v {
			return true
		}
		if input[next.row][next.col] != '#' && next != obstacle {
			point = next
			visited[visitor{point, dir}] = true
		} else {
			dir.TurnRight()
		}
	}
	return false
}

func day06_get_path(input []string, point Point, dir Point) (int, []visitor) {
	visited := map[Point]bool{point: false}
	path := []visitor{}
	for {
		next := point.Add(dir)
		if !InsideGrid(input, next) {
			break
		}
		if input[next.row][next.col] != '#' {
			path = append(path, visitor{next, dir})
			point = next
			visited[point] = true
		} else {
			dir.TurnRight()
		}
	}
	return len(visited), path
}
