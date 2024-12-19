package main

import (
	"fmt"
	"strings"
)

func day19() {
	// sample := []string{
	// 	"r, wr, b, g, bwu, rb, gb, br",
	// 	"",
	// 	"brwrr",
	// 	"bggr",
	// 	"gbbr",
	// 	"rrbgbr",
	// 	"ubwu",
	// 	"bwurrg",
	// 	"brgr",
	// 	"bbrgwb",
	// }
	input := ReadLines("input/day19.txt")
	patterns, designs := day19_parse(input)

	possible := []string{}
	for _, design := range designs {
		if day19_match(patterns, design) {
			possible = append(possible, design)
		}
	}

	fmt.Println("Day 19 Part 01:", len(possible))

	count := 0
	for _, design := range possible {
		count += day19_count_combinations(patterns, design, make(map[string]int))
	}

	fmt.Println("Day 19 Part 02:", count)
}

func day19_count_combinations(patterns []string, design string, cache map[string]int) int {
	if design == "" {
		return 1
	}
	if count, exists := cache[design]; exists {
		return count
	}

	total := 0
	for _, pattern := range patterns {
		if len(pattern) > len(design) {
			continue
		}
		if pattern == design[:len(pattern)] {
			total += day19_count_combinations(patterns, design[len(pattern):], cache)
		}
	}
	cache[design] = total
	return total
}

func day19_match(patterns []string, design string) bool {
	for _, pattern := range patterns {
		if pattern == design {
			return true
		}
		if len(pattern) > len(design) {
			continue
		}
		if pattern == design[:len(pattern)] && day19_match(patterns, design[len(pattern):]) {
			return true
		}
	}
	return false
}

func day19_parse(input []string) ([]string, []string) {
	patterns := []string{}
	designs := []string{}
	for i, line := range input {
		if line == "" {
			continue
		}
		if i == 0 {
			patterns = append(patterns, strings.Split(line, ", ")...)
		} else {
			designs = append(designs, line)
		}
	}
	return patterns, designs
}
