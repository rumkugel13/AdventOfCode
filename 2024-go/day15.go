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

	// sample2 := []string{
	// 	"##########",
	// 	"#..O..O.O#",
	// 	"#......O.#",
	// 	"#.OO..O.O#",
	// 	"#..O@..O.#",
	// 	"#O#..O...#",
	// 	"#O..O..O.#",
	// 	"#.OO.O.OO#",
	// 	"#....O...#",
	// 	"##########",
	// 	"",
	// 	"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^",
	// 	"vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v",
	// 	"><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<",
	// 	"<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^",
	// 	"^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><",
	// 	"^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^",
	// 	">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^",
	// 	"<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>",
	// 	"^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>",
	// 	"v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
	// }

	// sample3 := []string{
	// 	"#######",
	// 	"#...#.#",
	// 	"#.....#",
	// 	"#..OO@#",
	// 	"#..O..#",
	// 	"#.....#",
	// 	"#######",
	// 	"",
	// 	"<vv<<^^<<^^",
	// }

	input := ReadLines("input/day15.txt")
	grid := day15_parse_grid(input)
	moves := day15_parse_moves(input)
	robot := FindInCharGrid(grid, '@')
	for _, move := range moves {
		if day15_move(grid, robot, byte(move)) {
			robot = day15_next(robot, byte(move))
		}
	}
	result := day15_gps_sum(grid)

	fmt.Println("Day 15 Part 01:", result)

	grid = day15_grid_stretch(day15_parse_grid(input))
	robot = FindInCharGrid(grid, '@')
	for _, move := range moves {
		if day15_can_move2(grid, robot, byte(move)) {
			day15_move2(grid, robot, byte(move))
			robot = day15_next(robot, byte(move))
		}
	}
	result = day15_gps_sum(grid)

	fmt.Println("Day 15 Part 02:", result)
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
			if cell == 'O' || cell == '[' {
				sum += r*100 + c
			}
		}
	}
	return sum
}

func day15_move2(grid [][]byte, object Point, move byte) {
	next := day15_next(object, move)
	if grid[next.row][next.col] == '.' {
		CharGridSwap(grid, object, next)
	} else if (move == '^' || move == 'v') && grid[next.row][next.col] == '[' {
		nextRight := next.Add(Right)
		day15_move2(grid, next, move)
		day15_move2(grid, nextRight, move)
		CharGridSwap(grid, object, next)
	} else if (move == '^' || move == 'v') && grid[next.row][next.col] == ']' {
		nextLeft := next.Add(Left)
		day15_move2(grid, next, move)
		day15_move2(grid, nextLeft, move)
		CharGridSwap(grid, object, next)
	} else if grid[next.row][next.col] != '#' {
		day15_move2(grid, next, move)
		CharGridSwap(grid, object, next)
	}
}

func day15_can_move2(grid [][]byte, object Point, move byte) bool {
	next := day15_next(object, move)
	if grid[next.row][next.col] == '.' {
		return true
	}
	if grid[next.row][next.col] == '#' {
		return false
	}
	if move == '^' || move == 'v' {
		if grid[next.row][next.col] == '[' {
			right := next.Add(Right)
			return day15_can_move2(grid, next, move) && day15_can_move2(grid, right, move)
		}
		if grid[next.row][next.col] == ']' {
			left := next.Add(Left)
			return day15_can_move2(grid, next, move) && day15_can_move2(grid, left, move)
		}
	}
	return day15_can_move2(grid, next, move)
}

func day15_move(grid [][]byte, object Point, move byte) bool {
	next := day15_next(object, move)
	if grid[next.row][next.col] == '.' || (grid[next.row][next.col] == 'O' && day15_move(grid, next, move)) {
		CharGridSwap(grid, object, next)
		return true
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

func day15_parse_moves(input []string) string {
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

func day15_parse_grid(input []string) [][]byte {
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
