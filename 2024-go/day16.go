package main

import (
	"fmt"
	"math"
)

func day16() {
	// sample := []string{
	// 	"###############",
	// 	"#.......#....E#",
	// 	"#.#.###.#.###.#",
	// 	"#.....#.#...#.#",
	// 	"#.###.#####.#.#",
	// 	"#.#.#.......#.#",
	// 	"#.#.#####.###.#",
	// 	"#...........#.#",
	// 	"###.#.#####.#.#",
	// 	"#...#.....#.#.#",
	// 	"#.#.#.###.#.#.#",
	// 	"#.....#...#.#.#",
	// 	"#.###.#.#.#.#.#",
	// 	"#S..#.....#...#",
	// 	"###############",
	// }
	// sample2 := []string{
	// 	"#################",
	// 	"#...#...#...#..E#",
	// 	"#.#.#.#.#.#.#.#.#",
	// 	"#.#.#.#...#...#.#",
	// 	"#.#.#.#.###.#.#.#",
	// 	"#...#.#.#.....#.#",
	// 	"#.#.#.#.#.#####.#",
	// 	"#.#...#.#.#.....#",
	// 	"#.#.#####.#.###.#",
	// 	"#.#.#.......#...#",
	// 	"#.#.###.#####.###",
	// 	"#.#.#...#.....#.#",
	// 	"#.#.#.#####.###.#",
	// 	"#.#.#.........#.#",
	// 	"#.#.#.#########.#",
	// 	"#S#.............#",
	// 	"#################",
	// }
	input := ReadLines("input/day16.txt")
	start := FindInGrid(input, 'S')
	end := FindInGrid(input, 'E')
	score := day16_score(input, start, end)

	fmt.Println("Day 16 Part 01:", score)
}

func day16_score(grid []string, start, end Point) int {
	type Visitor struct {
		point Point
		dir   Point
	}

	visited := map[Point]int{}
	v := Visitor{start, Right}
	queue := []Visitor{v}
	visited[start] = 0
	minScore := math.MaxInt64

	for len(queue) > 0 {
		visitor := queue[0]
		queue = queue[1:]
		score := visited[visitor.point]

		if visitor.point == end {
			if score < minScore {
				minScore = score
			}
		}

		forward := visitor.point.Add(visitor.dir)
		if v, ok := visited[forward]; GridChar(grid, forward) != '#' && (!ok || v > score+1) {
			visited[forward] = score + 1
			queue = append(queue, Visitor{forward, visitor.dir})
		}

		left := visitor.point.Add(visitor.dir.Left())
		if v, ok := visited[left]; GridChar(grid, left) != '#' && (!ok || v > score+1001) {
			visited[left] = score + 1 + 1000
			queue = append(queue, Visitor{left, visitor.dir.Left()})
		}

		right := visitor.point.Add(visitor.dir.Right())
		if v, ok := visited[right]; GridChar(grid, right) != '#' && (!ok || v > score+1001) {
			visited[right] = score + 1 + 1000
			queue = append(queue, Visitor{right, visitor.dir.Right()})
		}
	}
	return minScore
}
