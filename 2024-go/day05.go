package main

import (
	"fmt"
	"slices"
	"sort"
)

func day05() {
	input := ReadLines("input/day05.txt")
	rules, updates := day05_parse(input)

	sum := 0
	for _, update := range updates {
		if day05_valid_update(rules, update) {
			sum += update[len(update)/2]
		}
	}

	fmt.Println("Day 05 Part 01:", sum)

	sum = 0
	for _, update := range updates {
		if !day05_valid_update(rules, update) {
			sort.SliceStable(update, func(i, j int) bool {
				for _, rule := range rules {
					if update[i] == rule[0] && update[j] == rule[1] {
						return true
					}
				}
				return false
			})
			sum += update[len(update)/2]
		}
	}

	fmt.Println("Day 05 Part 02:", sum)
}

func day05_valid_update(rules [][]int, update []int) bool {
	for _, rule := range rules {
		i1, i2 := slices.Index(update, rule[0]), slices.Index(update, rule[1])
		if i1 == -1 || i2 == -1 {
			continue
		}
		if i1 > i2 {
			return false
		}
	}
	return true
}

func day05_parse(input []string) ([][]int, [][]int) {
	rules := [][]int{}
	updates := [][]int{}
	a := true
	for _, line := range input {
		if line == "" {
			a = false
			continue
		}
		if a {
			rules = append(rules, SepToIntArr(line, "|"))
		} else {
			updates = append(updates, CommaSepToIntArr(line))
		}
	}
	return rules, updates
}
