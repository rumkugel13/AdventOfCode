package main

import (
	"os"
	"slices"
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

	Up    = Point{-1, 0}
	Down  = Point{1, 0}
	Left  = Point{0, -1}
	Right = Point{0, 1}
)

func OppositeDir(dir Point) Point {
	return Point{-dir.row, -dir.col}
}

func NeighborsInGrid(grid []string, pos Point) []Point {
	neighbors := []Point{}
	for _, dir := range Directions {
		neighbor := pos.Add(dir)
		if InsideGrid(grid, neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func Neighbors(pos Point) []Point {
	neighbors := []Point{}
	for _, dir := range Directions {
		neighbor := pos.Add(dir)
		neighbors = append(neighbors, neighbor)
	}
	return neighbors
}

func Corners(pos Point) []Point {
	corners := []Point{
		pos.Add(Point{0, 0}),
		pos.Add(Point{0, 1}),
		pos.Add(Point{1, 0}),
		pos.Add(Point{1, 1}),
	}
	return corners
}

func ReadLines(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}

func ReadLine(file string) string {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.ReplaceAll(string(data), "\r\n", "\n")
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

func (point *Point) Sub(other Point) Point {
	return Point{point.row - other.row, point.col - other.col}
}

func (point Point) Left() Point {
	return Point{-point.col, point.row}
}

func (point Point) Right() Point {
	return Point{point.col, -point.row}
}

func (point Point) Opposite() Point {
	return Point{-point.row, -point.col}
}

func (point *Point) TurnLeft() {
	*point = Point{-point.col, point.row}
}

func (point *Point) TurnRight() {
	*point = Point{point.col, -point.row}
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

func FindInCharGrid(grid [][]byte, char byte) Point {
	for row, line := range grid {
		for col, c := range line {
			if c == char {
				return Point{row, col}
			}
		}
	}
	return Point{-1, -1}
}

func FindAllInGrid(grid []string, char byte) []Point {
	pois := []Point{}
	for row, line := range grid {
		for col, c := range line {
			if c == rune(char) {
				pois = append(pois, Point{row, col})
			}
		}
	}
	return pois
}

func POIsInGrid(grid []string, except []byte) []Point {
	pois := []Point{}
	for row, line := range grid {
		for col, c := range line {
			if !slices.Contains(except, byte(c)) {
				pois = append(pois, Point{row, col})
			}
		}
	}
	return pois
}

func Distance(start, end Point) int {
	return Abs(start.row-end.row) + Abs(start.col-end.col)
}

func InsideGrid(grid []string, pos Point) bool {
	return pos.col >= 0 && pos.col < len(grid[0]) && pos.row >= 0 && pos.row < len(grid)
}

func GridValue(grid []string, pos Point) int {
	return int(grid[pos.row][pos.col] - '0')
}

func GridChar(grid []string, pos Point) byte {
	return grid[pos.row][pos.col]
}

func CharGridSwap(grid [][]byte, a, b Point) {
	grid[a.row][a.col], grid[b.row][b.col] = grid[b.row][b.col], grid[a.row][a.col]
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

func SpaceSepToIntArr(line string) []int {
	data := strings.Fields(line)
	result := make([]int, len(data))
	for i, val := range data {
		num, _ := strconv.Atoi(strings.TrimSpace(val))
		result[i] = num
	}
	return result
}

func SepToIntArr(line, sep string) []int {
	data := strings.Split(line, sep)
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
