package main

import "fmt"

func day15() {
	// sample := []string{
	// 	"########",
	// 	"#..O.O.#",
	// 	"##@.O..#",
	// 	"#...O..#",
	// 	"#.#.O..#",
	// 	"#...O..#",
	// 	"#......#",
	// 	"########",
	// 	"",
	// 	"<^^>>>vv<v>>v<<",
	// }

	input := ReadLines("input/day15.txt")
	grid := day15_grid(input)
	moves := day15_moves(input)
	grid = day15_move(grid, moves)
	result := day15_gps_sum(grid)

	fmt.Println("Day 15 Part 01:", result)

	grid = day15_grid_stretch(grid)

	fmt.Println("Day 15 Part 02:", "Not implemented yet")
}

func day15_grid_stretch(grid [][]byte) [][]byte {
	result := [][]byte{}
	for _, row := range grid {
		stretchedRow := []byte{}
		for _, cell := range row {
			switch cell {
			case '.':
				stretchedRow = append(stretchedRow, '.', '.')
			case '#':
				stretchedRow = append(stretchedRow, '#', '#')
			case 'O':
				stretchedRow = append(stretchedRow, '[', ']')
			case '@':
				stretchedRow = append(stretchedRow, '@', '.')
			}
		}
		result = append(result, stretchedRow)
	}
	return result
}

func day15_gps_sum(grid [][]byte) int {
	sum := 0
	for r, row := range grid {
		for c, cell := range row {
			if cell == 'O' {
				sum += r*100 + c
			}
		}
	}
	return sum
}

func day15_move(grid [][]byte, moves string) [][]byte {
	point := FindInCharGrid(grid, '@')
	for _, move := range moves {
		if day15_can_push(grid, point, byte(move)) {
			currentChar := grid[point.row][point.col]
			grid[point.row][point.col] = '.'

			next := day15_next(point, byte(move))
			for grid[next.row][next.col] != '.' {
				nextChar := grid[next.row][next.col]
				grid[next.row][next.col] = currentChar
				next = day15_next(next, byte(move))
				currentChar = nextChar
			}
			grid[next.row][next.col] = currentChar
			point = day15_next(point, byte(move))
		}
		// day15_print(grid)
	}
	return grid
}

func day15_print(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func day15_can_push(grid [][]byte, point Point, move byte) bool {
	next := day15_next(point, move)
	if grid[next.row][next.col] == 'O' {
		canPush := day15_can_push(grid, next, move)
		return canPush
	}
	return grid[next.row][next.col] == '.'
}

func day15_next(point Point, move byte) Point {
	var next Point
	switch move {
	case '>':
		next = Point{point.row, point.col + 1}
	case '<':
		next = Point{point.row, point.col - 1}
	case '^':
		next = Point{point.row - 1, point.col}
	case 'v':
		next = Point{point.row + 1, point.col}
	}
	return next
}

func day15_moves(input []string) string {
	moves := ""
	add := false
	for _, line := range input {
		if line == "" {
			add = true
		} else if add {
			moves += line
		}
	}
	return moves
}

func day15_grid(input []string) [][]byte {
	grid := [][]byte{}
	for _, line := range input {
		if line == "" {
			break
		} else {
			grid = append(grid, []byte(line))
		}
	}
	return grid
}
