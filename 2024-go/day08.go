package main

import "fmt"

func day08() {
	// sample := []string{
	// 	"............",
	// 	"........0...",
	// 	".....0......",
	// 	".......0....",
	// 	"....0.......",
	// 	"......A.....",
	// 	"............",
	// 	"............",
	// 	"........A...",
	// 	".........A..",
	// 	"............",
	// 	"............",
	// }
	input := ReadLines("input/day08.txt")
	pois := POIsInGrid(input, []byte{'.'})

	frequencies := map[byte][]Point{}
	for _, poi := range pois {
		frequencies[input[poi.row][poi.col]] = append(frequencies[input[poi.row][poi.col]], poi)
	}

	antinodes := day08_antinodes(frequencies, len(input[0]), len(input), false)
	fmt.Println("Day 08 Part 01:", len(antinodes))

	antinodes2 := day08_antinodes(frequencies, len(input[0]), len(input), true)
	fmt.Println("Day 08 Part 02:", len(antinodes2))
}

func day08_antinodes(frequencies map[byte][]Point, width, height int, part2 bool) map[Point]bool {
	antinodes := map[Point]bool{}
	for _, antennas := range frequencies {
		for m, a1 := range antennas {
			for n, a2 := range antennas {
				if m == n {
					continue
				}
				diff := a2.Sub(a1)
				node := a2.Add(diff)
				if !part2 {
					if node.col >= 0 && node.col < width && node.row >= 0 && node.row < height {
						antinodes[node] = true
					}
				} else {
					antinodes[a1] = true
					antinodes[a2] = true
					for node.col >= 0 && node.col < width && node.row >= 0 && node.row < height {
						antinodes[node] = true
						node = node.Add(diff)
					}
				}
			}
		}
	}
	return antinodes
}
