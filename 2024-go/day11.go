package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day11() {
	// sample := "0 1 10 99 999"
	// sample2 := "125 17"
	input := ReadLine("input/day11.txt")
	stones := strings.Fields(input)

	for range 25 {
		stones = day11_blink(stones)
	}

	fmt.Println("Day 11 Part 01:", len(stones))

	stones2 := SpaceSepToIntArr(input)
	stone_map := map[int]int{}
	for _, stone := range stones2 {
		stone_map[stone]++
	}
	for range 75 {
		stone_map = day11_blink2(stone_map)
	}
	sum := SumMap(stone_map)

	fmt.Println("Day 11 Part 02:", sum)
}

func day11_blink2(stones map[int]int) map[int]int {
	result := make(map[int]int, len(stones))
	for stone, amount := range stones {
		if stone == 0 {
			result[1] += amount
		} else {
			str := strconv.Itoa(stone)
			if (len(str) & 1) == 0 {
				v1 := str[0 : len(str)/2]
				v2 := str[len(str)/2:]
				n1, _ := strconv.Atoi(v1)
				n2, _ := strconv.Atoi(v2)
				result[n1] += amount
				result[n2] += amount
			} else {
				new_stone := stone * 2024
				result[new_stone] += amount
			}
		}
	}
	return result
}

func day11_blink(stones []string) []string {
	result := make([]string, 0, len(stones))
	for _, stone := range stones {
		if stone == "0" {
			result = append(result, "1")
		} else if (len(stone) & 1) == 0 {
			v1 := stone[0 : len(stone)/2]
			v2 := stone[len(stone)/2:]
			for len(v2) > 1 && v2[0] == '0' {
				v2 = v2[1:]
			}
			result = append(result, v1, v2)
		} else {
			num, _ := strconv.Atoi(stone)
			val := num * 2024
			str := strconv.Itoa(val)
			result = append(result, str)
		}
	}
	return result
}
