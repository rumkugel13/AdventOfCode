package main

import "fmt"

func day10() {
	// sample := []string{
	// 	"89010123",
	// 	"78121874",
	// 	"87430965",
	// 	"96549874",
	// 	"45678903",
	// 	"32019012",
	// 	"01329801",
	// 	"10456732",
	// }
	input := ReadLines("input/day10.txt")
	trailheads := FindAllInGrid(input, '0')
	sum := 0
	for _, trailhead := range trailheads {
		sum += day10_count_paths(input, trailhead)
	}

	fmt.Println("Day 10 Part 01:", sum)
	sum = 0
	for _, trailhead := range trailheads {
		sum += day10_count_distinct(input, trailhead)
	}
	fmt.Println("Day 10 Part 02:", sum)
}

func day10_count_distinct(grid []string, trailhead Point) int {
	paths := 0
	queue := []Point{trailhead}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		height := GridValue(grid, current)
		if height == 9 {
			paths++
			continue
		}

		for _, neighbor := range NeighborsInGrid(grid, current) {
			if GridValue(grid, neighbor) == height+1 {
				queue = append(queue, neighbor)
			}
		}
	}
	return paths
}

func day10_count_paths(grid []string, trailhead Point) int {
	visited := map[Point]bool{}
	visited[trailhead] = true
	paths := 0
	queue := []Point{trailhead}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		height := GridValue(grid, current)
		if height == 9 {
			paths++
			continue
		}

		for _, neighbor := range NeighborsInGrid(grid, current) {
			if _, ok := visited[neighbor]; !ok && GridValue(grid, neighbor) == height+1 {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return paths
}
