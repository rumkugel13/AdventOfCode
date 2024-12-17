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
	score, scores := day16_score(input, start, end)

	fmt.Println("Day 16 Part 01:", score)

	seats := day16_seats(scores, end, score)

	fmt.Println("Day 16 Part 02:", seats)
}

func day16_seats(scores map[Visitor]int, end Point, minscore int) int {
	type PointScore struct {
		point Point
		score int
	}

	seats := map[Point]bool{end: true}
	queue := []PointScore{{end, minscore}}
	for len(queue) > 0 {
		pointScore := queue[0]
		queue = queue[1:]

		for _, dir := range Directions {
			actualPoint := Visitor{pointScore.point, dir}

			if score, ok := scores[actualPoint]; ok {
				if score != pointScore.score {
					continue
				}

				prev := pointScore.point.Add(OppositeDir(dir))
				for _, dir2 := range Directions {
					prevVis := Visitor{prev, dir2}

					if val, found := scores[prevVis]; found {
						if (val == score-1) || (val == score-1001) {
							seats[prev] = true
							queue = append(queue, PointScore{prev, val})
						}
					}
				}
			}
		}
	}

	return len(seats)
}

type Visitor struct {
	point Point
	dir   Point
}

func day16_score(grid []string, start, end Point) (int, map[Visitor]int) {
	v := Visitor{start, Right}
	visited := map[Visitor]int{v: 0}
	queue := []Visitor{v}
	minScore := math.MaxInt64

	for len(queue) > 0 {
		visitor := queue[0]
		queue = queue[1:]
		score := visited[visitor]

		if visitor.point == end {
			if score < minScore {
				minScore = score
			}
			continue
		}

		dirs := []Point{visitor.dir, visitor.dir.Left(), visitor.dir.Right(), OppositeDir(visitor.dir)}
		newScores := []int{score + 1, score + 1001, score + 1001, score + 2001}
		for i, dir := range dirs {
			next := visitor.point.Add(dir)
			if GridChar(grid, next) == '#' {
				continue
			}
			nextVisitor := Visitor{next, dir}
			if v, ok := visited[nextVisitor]; !ok || v > newScores[i] {
				visited[nextVisitor] = newScores[i]
				queue = append(queue, Visitor{next, dir})
			}
		}
	}
	return minScore, visited
}
