package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day03() {
	input := ReadLine("input/day03.txt")
	// sample := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	// input := sample

	re := regexp.MustCompile(`mul\([0-9]*,[0-9]*\)`)
	matches := re.FindAllString(input, -1)

	sum := 0
	for _, match := range matches {
		nums := strings.Split(match[4:len(match)-1], ",")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		sum += num1 * num2
	}

	fmt.Println("Day 03 Part 01:", sum)

	// sample := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	// input = sample

	sum = 0
	do := true
	for i := 0; i < len(input); i++ {
		if i+4 < len(input) && input[i:i+4] == "do()" {
			do = true
			i += 3
		} else if i+7 < len(input) && input[i:i+7] == "don't()" {
			do = false
			i += 6
		} else if do && i+4 < len(input) && input[i:i+4] == "mul(" {
			match := re.FindStringIndex(input[i:])
			if match != nil {
				nums := strings.Split(input[i+match[0]+4:i+match[1]-1], ",")
				num1, _ := strconv.Atoi(nums[0])
				num2, _ := strconv.Atoi(nums[1])
				sum += num1 * num2
				i += match[1] - 1
			}
		}
	}

	fmt.Println("Day 03 Part 02:", sum)
}
