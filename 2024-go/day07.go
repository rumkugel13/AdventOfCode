package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day07() {
	// sample := []string {
	// 	"190: 10 19",
	// 	"3267: 81 40 27",
	// 	"83: 17 5",
	// 	"156: 15 6",
	// 	"7290: 6 8 6 15",
	// 	"161011: 16 10 13",
	// 	"192: 17 8 14",
	// 	"21037: 9 7 18 13",
	// 	"292: 11 6 16 20",
	// }

	input := ReadLines("input/day07.txt")
	data := day07_parse(input)
	sum := 0
	sum2 := 0
	for test, nums := range data {
		sum += day07_valid_equation(test, nums, false)
		sum2 += day07_valid_equation(test, nums, true)
	}

	fmt.Println("Day 07 Part 01:", sum)
	fmt.Println("Day 07 Part 02:", sum2)
}

func day07_valid_equation(testValue int, nums []int, part2 bool) int {
	results := day07_equation(testValue, nums, part2)
	for _, n := range results {
		if n == testValue {
			return testValue
		}
	}
	return 0
}

func day07_equation(testValue int, nums []int, part2 bool) []int {
	result := []int{}
	if len(nums) == 1 {
		result = append(result, nums[0])
		return result
	}

	next := day07_equation(testValue, nums[:len(nums)-1], part2)
	for _, n := range next {
		result = append(result, n+nums[len(nums)-1])
		result = append(result, n*nums[len(nums)-1])
		if part2 {
			concatenated, _ := strconv.Atoi(strconv.Itoa(n) + strconv.Itoa(nums[len(nums)-1]))
			result = append(result, concatenated)
		}
	}

	return result
}

func day07_parse(input []string) map[int][]int {
	data := make(map[int][]int)
	for _, line := range input {
		parts := strings.Split(line, ": ")
		num, _ := strconv.Atoi(parts[0])
		if _, ok := data[num]; ok {
			panic("Duplicate key")
		}
		data[num] = SpaceSepToIntArr(parts[1])
	}
	return data
}
