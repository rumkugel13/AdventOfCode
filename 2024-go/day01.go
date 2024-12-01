package main

import "fmt"
import "strings"
import "strconv"
import "sort"

func day01() {
	lines := ReadLines("input/day01.txt")
	l1, l2 := day01_parse(lines)
	
	sum := 0
	for i := 0; i < len(l1); i++ {
		sum += Abs(l1[i] - l2[i])
	}

	fmt.Println("Day 01 Part 01:", sum)

	amounts := map[int]int{}
	sum = 0
	for _, v1 := range l1 {
		amount := 0
		for _, v2 := range l2 {
			if v1 == v2 {
				amount++
			}
		}
		if _, ok := amounts[v1]; !ok {
			amounts[v1] = amount
		}
		sum += amounts[v1] * v1
	}

	fmt.Println("Day 01 Part 02:", sum)
}

func day01_parse(lines []string) ([]int, []int) {
	l1, l2 := []int{}, []int{}
	for _, line := range lines {
		fields := strings.Fields(line)
		v1, _ := strconv.Atoi(fields[0])
		v2, _ := strconv.Atoi(fields[1])
		l1 = append(l1, v1)
		l2 = append(l2, v2)
	}

	sort.Slice(l1, func(i, j int) bool { return l1[i] < l1[j] })
	sort.Slice(l2, func(i, j int) bool { return l2[i] < l2[j] })
	return l1, l2
}
