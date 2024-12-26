package main

import "fmt"

func day25() {
	// sample := []string{
	// 	"#####",
	// 	".####",
	// 	".####",
	// 	".####",
	// 	".#.#.",
	// 	".#...",
	// 	".....",
	// 	"",
	// 	"#####",
	// 	"##.##",
	// 	".#.##",
	// 	"...##",
	// 	"...#.",
	// 	"...#.",
	// 	".....",
	// 	"",
	// 	".....",
	// 	"#....",
	// 	"#....",
	// 	"#...#",
	// 	"#.#.#",
	// 	"#.###",
	// 	"#####",
	// 	"",
	// 	".....",
	// 	".....",
	// 	"#.#..",
	// 	"###..",
	// 	"###.#",
	// 	"###.#",
	// 	"#####",
	// 	"",
	// 	".....",
	// 	".....",
	// 	".....",
	// 	"#....",
	// 	"#.#..",
	// 	"#.#.#",
	// 	"#####",
	// }
	input := ReadLines("input/day25.txt")

	sum := 0
	locks, keys := day25_parse(input)
	for _, lock := range locks {
		for _, key := range keys {
			matches := true
			for i := range len(lock) {
				if lock[i]+key[i] > 5 {
					matches = false
					break
				}
			}
			if matches {
				sum++
			}
		}
	}

	fmt.Println("Day 25 Part 01:", sum)
}

func day25_parse(input []string) ([][5]int, [][5]int) {
	locks := [][5]int{}
	keys := [][5]int{}
	for i := 0; i < len(input); i += 8 {
		heights := [5]int{-1, -1, -1, -1, -1}
		for j := 0; j < 7; j++ {
			for col := 0; col < len(heights); col++ {
				if input[i+j][col] == '#' {
					heights[col]++
				}
			}
		}
		if input[i][0] == '#' {
			locks = append(locks, heights)
		} else {
			keys = append(keys, heights)
		}
	}
	return locks, keys
}
