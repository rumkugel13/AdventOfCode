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
	robot := FindInCharGrid(grid, '@')
	for _, move := range moves {
		day15_move(grid, &robot, byte(move))
	}
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

func day15_move(grid [][]byte, robot *Point, move byte) {
	if day15_can_push(grid, *robot, move) {
		currentChar := grid[robot.row][robot.col]
		grid[robot.row][robot.col] = '.'

		next := day15_next(*robot, move)
		for grid[next.row][next.col] != '.' {
			nextChar := grid[next.row][next.col]
			grid[next.row][next.col] = currentChar
			next = day15_next(next, move)
			currentChar = nextChar
		}
		grid[next.row][next.col] = currentChar
		*robot = day15_next(*robot, move)
	}
	// day15_print(grid)
}

func day15_print(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func day15_can_push(grid [][]byte, point Point, move byte) bool {
	next := day15_next(point, move)
	for grid[next.row][next.col] != '#' {
		if grid[next.row][next.col] == '.' {
			return true
		}
		next = day15_next(next, move)
	}

	return false
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
