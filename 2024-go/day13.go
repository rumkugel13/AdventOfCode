package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day13() {
	// sample := []string{
	// 	"Button A: X+94, Y+34",
	// 	"Button B: X+22, Y+67",
	// 	"Prize: X=8400, Y=5400",
	// 	"",
	// 	"Button A: X+26, Y+66",
	// 	"Button B: X+67, Y+21",
	// 	"Prize: X=12748, Y=12176",
	// 	"",
	// 	"Button A: X+17, Y+86",
	// 	"Button B: X+84, Y+37",
	// 	"Prize: X=7870, Y=6450",
	// 	"",
	// 	"Button A: X+69, Y+23",
	// 	"Button B: X+27, Y+71",
	// 	"Prize: X=18641, Y=10279",
	// }
	input := ReadLines("input/day13.txt")
	machines := day13_parse_machines(input)
	sum := 0
	for _, machine := range machines {
		sum += day13_tokens(machine, 0)
	}

	fmt.Println("Day 13 Part 01:", sum)

	sum = 0
	for _, machine := range machines {
		sum += day13_tokens(machine, 10000000000000)
	}

	fmt.Println("Day 13 Part 02:", sum)
}

func day13_tokens(machine Machine, offset int) int {
	a11 := machine.ButtonA.col
	a12 := machine.ButtonB.col
	a21 := machine.ButtonA.row
	a22 := machine.ButtonB.row
	n1 := machine.Prize.col + offset
	n2 := machine.Prize.row + offset

	determinant := a11*a22 - a12*a21
	if determinant == 0 {
		return 0
	}

	anum, bnum := (n1*a22 - a12*n2), (n2*a11 - a21*n1)
	a, arem := anum/determinant, anum%determinant
	b, brem := bnum/determinant, bnum%determinant

	if arem != 0 || brem != 0 {
		return 0
	}

	return a*3 + b
}

func day13_parse_machines(input []string) []Machine {
	machines := []Machine{}
	for line := 0; line < len(input); line += 4 {
		buttonA := Point{}
		parts := strings.Split(input[line], ": ")
		parts = strings.Split(parts[1], ", ")
		buttonA.col, _ = strconv.Atoi(parts[0][2:])
		buttonA.row, _ = strconv.Atoi(parts[1][2:])

		buttonB := Point{}
		parts = strings.Split(input[line+1], ": ")
		parts = strings.Split(parts[1], ", ")
		buttonB.col, _ = strconv.Atoi(parts[0][2:])
		buttonB.row, _ = strconv.Atoi(parts[1][2:])

		prize := Point{}
		parts = strings.Split(input[line+2], ": ")
		parts = strings.Split(parts[1], ", ")
		prize.col, _ = strconv.Atoi(parts[0][2:])
		prize.row, _ = strconv.Atoi(parts[1][2:])
		machines = append(machines, Machine{buttonA, buttonB, prize})
	}
	return machines
}

type Machine struct {
	ButtonA Point
	ButtonB Point
	Prize   Point
}
