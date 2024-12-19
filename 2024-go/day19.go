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

	possible := 0
	for _, design := range designs {
		if day19_match(patterns, design) {
			possible++
		}
	}

	fmt.Println("Day 19 Part 01:", possible)

	fmt.Println("Day 19 Part 02:", "Not implemented yet")
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
