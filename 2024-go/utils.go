package main

import (
	"os"
	"strconv"
	"strings"
)

var (
	Directions = [4]Point{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
)

func ReadLines(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

type Point struct {
	row, col int
}

func (point *Point) Add(other Point) Point {
	return Point{point.row + other.row, point.col + other.col}
}

func FindInGrid(grid []string, char byte) Point {
	for row, line := range grid {
		for col, c := range line {
			if c == rune(char) {
				return Point{row, col}
			}
		}
	}
	return Point{-1, -1}
}

func Distance(start, end Point) int {
	return Abs(start.row-end.row) + Abs(start.col-end.col)
}

func InsideGrid(grid []string, pos Point) bool {
	return pos.col >= 0 && pos.col < len(grid[0]) && pos.row >= 0 && pos.row < len(grid)
}

func Duplicate[T any](grid [][]T) [][]T {
	duplicate := make([][]T, len(grid))
	for i := range grid {
		duplicate[i] = make([]T, len(grid[i]))
		copy(duplicate[i], grid[i])
	}
	return duplicate
}

func DuplicateMap[T comparable, U any](source map[T]U) map[T]U {
	duplicate := make(map[T]U)
	for key, value := range source {
		duplicate[key] = value
	}
	return duplicate
}

func SumMap[T comparable, U int | float32 | float64](source map[T]U) U {
	var sum U
	for _, value := range source {
		sum += value
	}
	return sum
}

func SumSlice(source []int) int {
	var sum int
	for _, value := range source {
		sum += value
	}
	return sum
}

func Mod(a, n int) int {
	return ((a % n) + n) % n
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func CommaSepToIntArr(line string) []int {
	data := strings.Split(line, ",")
	result := make([]int, len(data))
	for i, val := range data {
		num, _ := strconv.Atoi(strings.TrimSpace(val))
		result[i] = num
	}
	return result
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func LCMSlice(arr []int) int {
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = LCM(result, arr[i])
	}
	return result
}
