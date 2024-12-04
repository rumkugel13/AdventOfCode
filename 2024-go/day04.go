package main

import "fmt"

func day04() {
	input := ReadLines("input/day04.txt")
	result := day04_count_xmas(input)

	fmt.Println("Day 04 Part 01:", result)

	result = day04_count_mas(input)
	fmt.Println("Day 04 Part 02:", result)
}

func day04_count_mas(input []string) int {
	count := 0
	for r, line := range input {
		for c, ch := range line {
			if ch == 'A' {
				count += day04_count_mas_diag(input, r, c)
			}
		}
	}
	return count
}

func day04_count_mas_diag(input []string, r, c int) int {
	count := 0
	if r-1 >= 0 && r+1 < len(input) && c-1 >= 0 && c+1 < len(input[r]) {
		if ((input[r-1][c-1] == 'M' && input[r+1][c+1] == 'S') || (input[r-1][c-1] == 'S' && input[r+1][c+1] == 'M')) &&
			((input[r-1][c+1] == 'M' && input[r+1][c-1] == 'S') || (input[r-1][c+1] == 'S' && input[r+1][c-1] == 'M')) {
			count++
		}
	}
	return count
}

func day04_count_xmas(input []string) int {
	count := 0
	for r, line := range input {
		for c, ch := range line {
			if ch == 'X' {
				count += day04_count_hor(input, r, c)
				count += day04_count_vert(input, r, c)
				count += day04_count_diag(input, r, c)
			}
		}
	}
	return count
}

func day04_count_hor(input []string, r, c int) int {
	count := 0
	if c+3 < len(input[r]) && input[r][c+1] == 'M' && input[r][c+2] == 'A' && input[r][c+3] == 'S' {
		count++
	}
	if c-3 >= 0 && input[r][c-1] == 'M' && input[r][c-2] == 'A' && input[r][c-3] == 'S' {
		count++
	}
	return count
}

func day04_count_vert(input []string, r, c int) int {
	count := 0
	if r+3 < len(input) && input[r+1][c] == 'M' && input[r+2][c] == 'A' && input[r+3][c] == 'S' {
		count++
	}
	if r-3 >= 0 && input[r-1][c] == 'M' && input[r-2][c] == 'A' && input[r-3][c] == 'S' {
		count++
	}
	return count
}

func day04_count_diag(input []string, r, c int) int {
	count := 0
	if r+3 < len(input) && c+3 < len(input[r]) && input[r+1][c+1] == 'M' && input[r+2][c+2] == 'A' && input[r+3][c+3] == 'S' {
		count++
	}
	if r-3 >= 0 && c-3 >= 0 && input[r-1][c-1] == 'M' && input[r-2][c-2] == 'A' && input[r-3][c-3] == 'S' {
		count++
	}
	if r+3 < len(input) && c-3 >= 0 && input[r+1][c-1] == 'M' && input[r+2][c-2] == 'A' && input[r+3][c-3] == 'S' {
		count++
	}
	if r-3 >= 0 && c+3 < len(input[r]) && input[r-1][c+1] == 'M' && input[r-2][c+2] == 'A' && input[r-3][c+3] == 'S' {
		count++
	}
	return count
}
