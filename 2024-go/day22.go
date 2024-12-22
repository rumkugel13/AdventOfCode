package main

import (
	"fmt"
	"strconv"
)

func day22() {
	// sample := []string{
	// 	"1",
	// 	"10",
	// 	"100",
	// 	"2024",
	// }
	input := ReadLines("input/day22.txt")
	numbers := day22_numbers(input)
	for range 2000 {
		for i, num := range numbers {
			numbers[i] = day22_next_number(num)
		}
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	fmt.Println("Day 22 Part 01:", sum)
}

func day22_next_number(num int) int {
	num = ((num * 64) ^ num) % 16777216
	num = ((num / 32) ^ num) % 16777216
	num = ((num * 2048) ^ num) % 16777216
	return num
}

func day22_numbers(input []string) []int {
	numbers := make([]int, len(input))
	for i, line := range input {
		numbers[i], _ = strconv.Atoi(line)
	}
	return numbers
}
