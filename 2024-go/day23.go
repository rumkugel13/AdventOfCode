package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

func day23() {
	// sample := []string{
	// 	"kh-tc",
	// 	"qp-kh",
	// 	"de-cg",
	// 	"ka-co",
	// 	"yn-aq",
	// 	"qp-ub",
	// 	"cg-tb",
	// 	"vc-aq",
	// 	"tb-ka",
	// 	"wh-tc",
	// 	"yn-cg",
	// 	"kh-ub",
	// 	"ta-co",
	// 	"de-co",
	// 	"tc-td",
	// 	"tb-wq",
	// 	"wh-td",
	// 	"ta-ka",
	// 	"td-qp",
	// 	"aq-cg",
	// 	"wq-ub",
	// 	"ub-vc",
	// 	"de-ta",
	// 	"wq-aq",
	// 	"wq-vc",
	// 	"wh-yn",
	// 	"ka-de",
	// 	"kh-ta",
	// 	"co-tc",
	// 	"wh-qp",
	// 	"tb-vc",
	// 	"td-yn",
	// }
	input := ReadLines("input/day23.txt")
	connections := day23_connections(input)
	sets := day23_sets_of_three(connections)

	count := 0
	for set := range sets {
		split := strings.Split(set, ",")
		for _, s := range split {
			if s[0] == 't' {
				count++
				break
			}
		}
	}

	fmt.Println("Day 23 Part 01:", count)

	sets = day23_sets(connections)
	longestSet := ""
	for set := range sets {
		if len(set) > len(longestSet) {
			longestSet = set
		}
	}

	fmt.Println("Day 23 Part 02:", longestSet)
}

func day23_sets(connections map[string][]string) map[string]bool {
	sets := map[string]bool{}
	for first, connection := range connections {
		set := map[string]bool{first: true}
		for i, second := range connection {
			add := true
			for j := i + 1; j < len(connection); j++ {
				third := connection[j]
				if !slices.Contains(connections[second], third) {
					add = false
					break
				}
			}
			if add {
				set[second] = true
			}
		}
		slice := slices.Collect(maps.Keys(set))
		slices.Sort(slice)
		sets[strings.Join(slice, ",")] = true
	}
	return sets
}

func day23_sets_of_three(connections map[string][]string) map[string]bool {
	sets := map[string]bool{}
	for first, connection := range connections {
		for i, second := range connection {
			for j := i + 1; j < len(connection); j++ {
				third := connection[j]
				if slices.Contains(connections[second], third) {
					set := []string{first, second, third}
					slices.Sort(set)
					sets[strings.Join(set, ",")] = true
				}
			}
		}
	}
	return sets
}

func day23_connections(input []string) map[string][]string {
	connections := make(map[string][]string)
	for _, line := range input {
		parts := strings.Split(line, "-")
		connections[parts[0]] = append(connections[parts[0]], parts[1])
		connections[parts[1]] = append(connections[parts[1]], parts[0])
	}
	return connections
}
