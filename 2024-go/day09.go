package main

import (
	"fmt"
	"slices"
)

func day09() {
	// sample := "2333133121414131402"
	input := ReadLine("input/day09.txt")
	blocks := day09_file_blocks(input)
	compact := day09_compact(blocks)
	checksum := day09_checksum(compact)

	fmt.Println("Day 09 Part 01:", checksum)

	blocks2 := day09_file_blocks(input)
	compact2 := day09_compact2(blocks2)
	checksum2 := day09_checksum(compact2)

	fmt.Println("Day 09 Part 02:", checksum2)
}

func day09_checksum(input []int) int {
	sum := 0
	for i, val := range input {
		if input[i] != -1 {
			sum += val * i
		}
	}
	return sum
}

func day09_compact2(input []int) []int {
	for id := input[len(input)-1]; id > 0; id-- {
		rstart := slices.Index(input, id)
		len_required := 0
		for i := rstart; i < len(input) && input[i] == id; i++ {
			len_required++
		}

		lstart, lend := 0, 0
		for (lend - lstart) < len_required {
			lstart = lend + slices.Index(input[lend:], -1)
			for i := lstart; i < len(input) && input[i] == -1; i++ {
				lend = i
			}
			lend++
		}

		if lstart > rstart {
			continue
		}

		for i := 0; i < len_required; i++ {
			input[i+lstart] = id
			input[i+rstart] = -1
		}
	}
	return input
}

func day09_compact(input []int) []int {
	for l, r := 0, len(input)-1; l < r; {
		for l < r && input[l] != -1 {
			l++
		}
		for l < r && input[r] == -1 {
			r--
		}
		if l < r {
			input[l], input[r] = input[r], input[l]
		}
	}
	return input
}

func day09_file_blocks(input string) []int {
	blocks := []int{}
	for i, id := 0, 0; i < len(input); i++ {
		amount := int(input[i] - '0')
		if (i & 1) == 0 {
			blocks = append(blocks, slices.Repeat([]int{id}, amount)...)
			id++
		} else {
			blocks = append(blocks, slices.Repeat([]int{-1}, amount)...)
		}
	}
	return blocks
}
