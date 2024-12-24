package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day24() {
	// sample := []string{
	// 	"x00: 1",
	// 	"x01: 1",
	// 	"x02: 1",
	// 	"y00: 0",
	// 	"y01: 1",
	// 	"y02: 0",
	// 	"",
	// 	"x00 AND y00 -> z00",
	// 	"x01 XOR y01 -> z01",
	// 	"x02 OR y02 -> z02",
	// }

	// sample2 := []string{
	// 	"x00: 1",
	// 	"x01: 0",
	// 	"x02: 1",
	// 	"x03: 1",
	// 	"x04: 0",
	// 	"y00: 1",
	// 	"y01: 1",
	// 	"y02: 1",
	// 	"y03: 1",
	// 	"y04: 1",
	// 	"",
	// 	"ntg XOR fgs -> mjb",
	// 	"y02 OR x01 -> tnw",
	// 	"kwq OR kpj -> z05",
	// 	"x00 OR x03 -> fst",
	// 	"tgd XOR rvg -> z01",
	// 	"vdt OR tnw -> bfw",
	// 	"bfw AND frj -> z10",
	// 	"ffh OR nrd -> bqk",
	// 	"y00 AND y03 -> djm",
	// 	"y03 OR y00 -> psh",
	// 	"bqk OR frj -> z08",
	// 	"tnw OR fst -> frj",
	// 	"gnj AND tgd -> z11",
	// 	"bfw XOR mjb -> z00",
	// 	"x03 OR x00 -> vdt",
	// 	"gnj AND wpb -> z02",
	// 	"x04 AND y00 -> kjc",
	// 	"djm OR pbm -> qhw",
	// 	"nrd AND vdt -> hwm",
	// 	"kjc AND fst -> rvg",
	// 	"y04 OR y02 -> fgs",
	// 	"y01 AND x02 -> pbm",
	// 	"ntg OR kjc -> kwq",
	// 	"psh XOR fgs -> tgd",
	// 	"qhw XOR tgd -> z09",
	// 	"pbm OR djm -> kpj",
	// 	"x03 XOR y03 -> ffh",
	// 	"x00 XOR y04 -> ntg",
	// 	"bfw OR bqk -> z06",
	// 	"nrd XOR fgs -> wpb",
	// 	"frj XOR qhw -> z04",
	// 	"bqk OR frj -> z07",
	// 	"y03 OR x01 -> nrd",
	// 	"hwm AND bqk -> z03",
	// 	"tgd XOR rvg -> z12",
	// 	"tnw OR pbm -> gnj",
	// }
	input := ReadLines("input/day24.txt")
	wires, gates := day24_parse(input)
	day24_propagate(wires, gates)

	sum := 0
	for wire, val := range wires {
		if wire[0] == 'z' && val {
			bit, _ := strconv.Atoi(string(wire[1:]))
			sum += 1 << bit
		}
	}

	fmt.Println("Day 24 Part 01:", sum)
}

func day24_propagate(wires map[string]bool, gates []Gate) {
	changed := true
	for changed {
		changed = false
		for _, gate := range gates {
			if _, found := wires[gate.res]; found {
				continue
			}
			op1, ok := wires[gate.op1]
			op2, ok2 := wires[gate.op2]
			if !ok || !ok2 {
				continue
			}
			switch gate.op {
			case "AND":
				wires[gate.res] = op1 && op2
			case "OR":
				wires[gate.res] = op1 || op2
			case "XOR":
				wires[gate.res] = op1 != op2
			}
			changed = true
		}
	}
}

type Gate struct {
	op1 string
	op2 string
	op  string
	res string
}

func day24_parse(input []string) (map[string]bool, []Gate) {
	wires := make(map[string]bool)
	gates := make([]Gate, 0)
	for _, line := range input {
		parts := strings.Split(line, " ")
		if len(parts) == 2 {
			val := false
			if parts[1] == "1" {
				val = true
			}
			wires[parts[0][:len(parts[0])-1]] = val
		} else if len(parts) == 5 {
			gate := Gate{parts[0], parts[2], parts[1], parts[4]}
			gates = append(gates, gate)
		}
	}
	return wires, gates
}
