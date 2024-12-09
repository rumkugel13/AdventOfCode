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
	for i := 0; i < len(input); i++ {
		if input[i] == -1 {
			continue
		}
		sum += int(input[i]) * i
	}
	return sum
}

func day09_compact2(input []int) []int {
	for id := input[len(input)-1]; id > 0; id-- {
		rstart := slices.Index(input, id)
		rend := rstart
		for i := rstart; i < len(input) && input[i] == id; i++ {
			rend = i
		}
		len_required := rend - rstart + 1

		lstart, lend := 0, 0
		freespace := 0
		for freespace < len_required && lend < len(input) {
			off := lend
			lstart = off + slices.Index(input[lend:], -1)
			for i := lstart; i < len(input) && input[i] == -1; i++ {
				lend = i
			}
			lend++
			freespace = lend - lstart
		}

		if freespace < len_required || lstart > rstart {
			continue
		}

		for i := 0; i < len_required; i++ {
			input[i+lstart] = id
		}
		for i := rstart; i <= rend; i++ {
			input[i] = -1
		}
	}
	return input
}

func day09_compact(input []int) []int {
	l := 0
	r := len(input) - 1
	for l < r {
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
	id := 0
	for i := 0; i < len(input); i++ {
		if (i & 1) == 0 {
			files := make([]int, input[i]-'0')
			for f := range len(files) {
				files[f] = int(id)
			}
			blocks = append(blocks, files...)
			id++
		} else {
			freespace := make([]int, input[i]-'0')
			for f := range len(freespace) {
				freespace[f] = -1
			}
			blocks = append(blocks, freespace...)
		}
	}
	return blocks
}
