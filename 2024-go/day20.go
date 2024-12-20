package main

import (
	"fmt"
)

func day20() {
	// sample := []string{
	// 	"###############",
	// 	"#...#...#.....#",
	// 	"#.#.#.#.#.###.#",
	// 	"#S#...#.#.#...#",
	// 	"#######.#.#.###",
	// 	"#######.#.#...#",
	// 	"#######.#.###.#",
	// 	"###..E#...#...#",
	// 	"###.#######.###",
	// 	"#...###...#...#",
	// 	"#.#####.#.###.#",
	// 	"#.#...#.#.#...#",
	// 	"#.#.#.#.#.#.###",
	// 	"#...#...#...###",
	// 	"###############",
	// }
	input := ReadLines("input/day20.txt")
	start, end := FindInGrid(input, 'S'), FindInGrid(input, 'E')
	path := day20_path(input, start, end)
	cheats := day20_cheat_count(path, 2)

	fmt.Println("Day 20 Part 01:", cheats)

	cheats2 := day20_cheat_count(path, 20)

	fmt.Println("Day 20 Part 02:", cheats2)
}

type Step struct {
	point Point
	steps int
}

func day20_cheat_count(path []Step, maxDist int) int {
	cheats := 0
	for i, step := range path {
		for j := i + 101; j < len(path); j++ {
			next := path[j]
			dist := Distance(step.point, next.point)
			if dist <= maxDist && (next.steps-step.steps-dist) >= 100 {
				cheats++
			}
		}
	}
	return cheats
}

func day20_path(input []string, start, end Point) []Step {
	queue := []Point{start}
	visited := map[Point]int{start: 0}
	path := []Step{{start, 0}}
	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		steps := visited[point]
		if point == end {
			break
		}
		for _, next := range Neighbors(point) {
			if _, ok := visited[next]; !ok && input[next.row][next.col] != '#' {
				visited[next] = steps + 1
				path = append(path, Step{next, steps + 1})
				queue = append(queue, next)
			}
		}
	}
	return path
}
