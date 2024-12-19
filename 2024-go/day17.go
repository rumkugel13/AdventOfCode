package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func day17() {
	// sample := []string{
	// 	"Register A: 729",
	// 	"Register B: 0",
	// 	"Register C: 0",
	// 	"",
	// 	"Program: 0,1,5,4,3,0",
	// }
	// sample2 := []string{
	// 	"Register A: 2024",
	// 	"Register B: 0",
	// 	"Register C: 0",
	// 	"",
	// 	"Program: 0,3,5,4,3,0",
	// }
	input := ReadLines("input/day17.txt")
	computer := day17_parse_computer(input)
	day17_run_computer(&computer)
	result := strconv.Itoa(computer.out[0])
	for i := 1; i < len(computer.out); i++ {
		result += "," + strconv.Itoa(computer.out[i])
	}

	fmt.Println("Day 17 Part 01:", result)

	computer = day17_parse_computer(input)
	possibleInitialValues := map[int]bool{0: true}

	// adapted from https://github.com/derailed-dash/Advent-of-Code/blob/master/src/AoC_2024/Dazbo's_Advent_of_Code_2024.ipynb#Day-17-Part-2
	for position := len(computer.prog) - 1; position >= 0; position-- {
		targetInstruction := computer.prog[position]
		nextPossibleValues := make(map[int]bool)

		for currentValue := range possibleInitialValues {
			shiftedValue := currentValue * 8

			for testValue := shiftedValue; testValue < shiftedValue+8; testValue++ {
				computer.pc = 0
				computer.regs["A"] = testValue
				computer.regs["B"] = 0
				computer.regs["C"] = 0
				computer.out = []int{}
				day17_run_computer(&computer)

				if len(computer.out) > 0 && computer.out[0] == targetInstruction {
					nextPossibleValues[testValue] = true
				}
			}
		}

		possibleInitialValues = nextPossibleValues
	}

	minInitialValue := math.MaxInt64
	for value := range possibleInitialValues {
		if value < minInitialValue {
			minInitialValue = value
		}
	}

	fmt.Println("Day 17 Part 02:", minInitialValue)
}

func day17_run_computer(computer *Computer) {
	for ; computer.pc < len(computer.prog); computer.pc += 2 {
		instruction := computer.prog[computer.pc]
		operand := computer.prog[computer.pc+1]
		comboRegister := 'A' + operand - 4
		comboOperand := operand
		if comboRegister >= 'A' && comboRegister <= 'C' {
			comboOperand = computer.regs[string(byte(comboRegister))]
		}

		switch instruction {
		case 0:
			numerator := computer.regs["A"]
			denominator := 1 << comboOperand
			computer.regs["A"] = (numerator / denominator)
		case 1:
			val1 := computer.regs["B"]
			val2 := operand
			computer.regs["B"] = val1 ^ val2
		case 2:
			val := comboOperand & 0b111
			computer.regs["B"] = val
		case 3:
			if computer.regs["A"] != 0 {
				computer.pc = operand - 2
			}
		case 4:
			val1 := computer.regs["B"]
			val2 := computer.regs["C"]
			computer.regs["B"] = val1 ^ val2
		case 5:
			computer.out = append(computer.out, comboOperand&0b111)
		case 6:
			numerator := computer.regs["A"]
			denominator := 1 << comboOperand
			computer.regs["B"] = numerator / denominator
		case 7:
			numerator := computer.regs["A"]
			denominator := 1 << comboOperand
			computer.regs["C"] = numerator / denominator
		}
	}
}

func day17_parse_computer(input []string) Computer {
	computer := Computer{
		regs: map[string]int{},
		prog: []int{},
		pc:   0,
		out:  []int{},
	}

	for _, line := range input {
		parts := strings.Split(line, " ")
		if parts[0] == "Register" {
			num, _ := strconv.Atoi(parts[2])
			computer.regs[parts[1][:1]] = num
		} else if parts[0] == "Program:" {
			computer.prog = CommaSepToIntArr(parts[1])
		}
	}
	return computer
}

type Computer struct {
	regs map[string]int
	prog []int
	pc   int
	out  []int
}
