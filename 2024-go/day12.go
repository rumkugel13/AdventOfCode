package main

import (
	"fmt"
	"slices"
)

func day12() {
	// sample := []string{
	// 	"AAAA",
	// 	"BBCD",
	// 	"BBCC",
	// 	"EEEC",
	// }
	// sample2 := []string{
	// 	"OOOOO",
	// 	"OXOXO",
	// 	"OOOOO",
	// 	"OXOXO",
	// 	"OOOOO",
	// }
	// sample3 := []string{
	// 	"RRRRIICCFF",
	// 	"RRRRIICCCF",
	// 	"VVRRRCCFFF",
	// 	"VVRCCCJFFF",
	// 	"VVVVCJJCFE",
	// 	"VVIVCCJJEE",
	// 	"VVIIICJJEE",
	// 	"MIIIIIJJEE",
	// 	"MIIISIJEEE",
	// 	"MMMISSJEEE",
	// }
	// sample4 := []string{
	// 	"EEEEE",
	// 	"EXXXX",
	// 	"EEEEE",
	// 	"EXXXX",
	// 	"EEEEE",
	// }
	// sample5 := []string{
	// 	"AAAAAA",
	// 	"AAABBA",
	// 	"AAABBA",
	// 	"ABBAAA",
	// 	"ABBAAA",
	// 	"AAAAAA",
	// }
	input := ReadLines("input/day12.txt")
	regions := day12_parse_regions(input)
	sum := 0
	for _, region := range regions {
		perimeter := day12_perimeter(region)
		sum += perimeter * len(region)
	}

	fmt.Println("Day 12 Part 01:", sum)

	sum = 0
	for _, region := range regions {
		sides := day12_sides(region)
		sum += sides * len(region)
	}

	fmt.Println("Day 12 Part 02:", sum)
}

func day12_sides(region []Point) int {
	corners := map[Point][]Point{}
	for _, point := range region {
		cornersOnPoint := Corners(point)
		for _, corner := range cornersOnPoint {
			corners[corner] = append(corners[corner], point)
		}
	}

	sum := 0
	for _, corner := range corners {
		if (len(corner) & 1) == 1 {
			sum++
		} else if len(corner) == 2 {
			neighbors := Neighbors(corner[0])
			if !slices.Contains(neighbors, corner[1]) {
				sum += 2
			}
		}
	}
	return sum
}

func day12_perimeter(region []Point) int {
	perimeter := 0
	for _, point := range region {
		for _, neighbor := range Neighbors(point) {
			if !slices.Contains(region, neighbor) {
				perimeter++
			}
		}
	}
	return perimeter
}

func day12_parse_regions(grid []string) [][]Point {
	regions := [][]Point{}
	visited := make(map[Point]bool)

	for r, row := range grid {
		for c, char := range row {
			if _, v := visited[Point{r, c}]; v {
				continue
			}
			pos := Point{r, c}
			queue := []Point{pos}
			region := []Point{pos}
			currentChar := byte(char)
			visited[pos] = true

			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]

				for _, neighbor := range NeighborsInGrid(grid, current) {
					if _, v := visited[neighbor]; v {
						continue
					}

					if GridChar(grid, neighbor) == currentChar {
						visited[neighbor] = true
						queue = append(queue, neighbor)
						region = append(region, neighbor)
					}
				}
			}

			regions = append(regions, region)
		}
	}
	return regions
}
