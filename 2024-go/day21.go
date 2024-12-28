package main

import (
	"fmt"
	// "slices"
	"strconv"
)

func day21() {
	sample := []string{
		"029A",
		"980A",
		"179A",
		"456A",
		"379A",
	}
	input := sample //ReadLines("input/day21.txt")
	numericKeypad := []string{
		"789",
		"456",
		"123",
		" 0A",
	}
	dirKeypad := []string{
		" ^A",
		"<v>",
	}

	sum := 0
	for _,code := range input {
		sequence := day21_sequence_needed(code, numericKeypad)
		sequence = day21_sequence_needed(sequence, dirKeypad)
		sequence = day21_sequence_needed(sequence, dirKeypad)
		num, _ := strconv.Atoi(code[:len(code)-1])
		// fmt.Println(len(sequence), num, sequence)
		complexity := len(sequence) * num
		sum += complexity
	}

	fmt.Println("Day 21 Part 01:", sum)
	fmt.Println("Day 21 Part 02:", "Not immplemented yet")
}

func day21_sequence_needed(code string, keypad []string) string {
	start := FindInGrid(keypad, 'A')
	sequence := ""
	for _, char := range code {
		end := FindInGrid(keypad, byte(char))
		sequence += day21_next_position(start, end, keypad)
		start = end
	}
	return sequence
}

func day21_next_position(start Point, end Point, keypad []string) string {
	type queueItem struct {
		point    Point
		sequence string
	}
	// aKey := FindInGrid(keypad, 'A')

	queue := []queueItem{{start, ""}}
	seen := map[Point]bool{start: true}
	minSeq := ""
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if item.point == end {
			if minSeq == "" || len(item.sequence) < len(minSeq) {
				minSeq = item.sequence
			}
		}

		// neighbors := NeighborsInGrid(keypad, item.point)
		// slices.SortFunc(neighbors, func(a, b Point) int {
		// 	return Distance(b, aKey) - Distance(a, aKey)
		// })

		for _, dir := range []Point{Left,Down,Right,Up} {
			next := item.point.Add(dir)
			if InsideGrid(keypad, next) && !seen[next] && keypad[next.row][next.col] != ' ' {
				seen[next] = true
				dir := Point{next.row - item.point.row, next.col - item.point.col}
				queue = append(queue, queueItem{next, item.sequence + string(day21_dir_to_char(dir))})
			}
		}
	}
	return minSeq + "A"
}

func day21_dir_to_char(dir Point) byte {
	switch dir {
	case Up:
		return '^'
	case Down:
		return 'v'
	case Left:
		return '<'
	case Right:
		return '>'
	}
	return ' '
}

func day21_char_to_dir(char byte) Point {
	switch char {
	case '^':
		return Up
	case 'v':
		return Down
	case '<':
		return Left
	case '>':
		return Right
	}
	return Point{}
}
