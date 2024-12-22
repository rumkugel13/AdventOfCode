package main

import (
	"fmt"
	"strconv"
)

func day22() {
	// sample := []string{
	// 	"1",
	// 	"10",
	// 	"100",
	// 	"2024",
	// }
	// sample2 := []string{
	// 	"1",
	// 	"2",
	// 	"3",
	// 	"2024",
	// }
	input := ReadLines("input/day22.txt")
	secrets := day22_secrets(input)
	for range 2000 {
		for i, num := range secrets {
			secrets[i] = day22_next_secret(num)
		}
	}

	sum := 0
	for _, num := range secrets {
		sum += num
	}

	fmt.Println("Day 22 Part 01:", sum)

	secrets = day22_secrets(input)
	allBananas := make([][]int, len(secrets))
	allChanges := make([][]int, len(secrets))
	for buyer := range secrets {
		currentSecret := secrets[buyer]
		allBananas[buyer] = make([]int, 2001)
		allBananas[buyer][0] = currentSecret % 10
		allChanges[buyer] = make([]int, 2000)
		for j := range 2000 {
			nextSecret := day22_next_secret(currentSecret)
			allBananas[buyer][j+1] = nextSecret % 10
			currentSecret = nextSecret
			allChanges[buyer][j] = allBananas[buyer][j+1] - allBananas[buyer][j]
		}
	}

	bananas := map[[4]int]int{}
	for buyer := range secrets {
		sequenceSeen := map[[4]int]bool{}
		for seqStart := 0; seqStart < 2000-3; seqStart++ {
			sequence := [4]int(allChanges[buyer][seqStart : seqStart+4])
			if sequenceSeen[sequence] {
				continue
			}
			bananas[sequence] += allBananas[buyer][seqStart+4]
			sequenceSeen[sequence] = true
		}
	}

	maxBananas := 0
	for _, bananas := range bananas {
		if bananas > maxBananas {
			maxBananas = bananas
		}
	}

	fmt.Println("Day 22 Part 02:", maxBananas)
}

func day22_next_secret(num int) int {
	num = ((num * 64) ^ num) % 16777216
	num = ((num / 32) ^ num) % 16777216
	num = ((num * 2048) ^ num) % 16777216
	return num
}

func day22_secrets(input []string) []int {
	numbers := make([]int, len(input))
	for i, line := range input {
		numbers[i], _ = strconv.Atoi(line)
	}
	return numbers
}
