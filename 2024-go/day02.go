package main

import (
	"fmt"
)

func day02() {
	input := ReadLines("input/day02.txt")
	// sample := []string{
	// 	"7 6 4 2 1",
	// 	"1 2 7 8 9",
	// 	"9 7 6 2 1",
	// 	"1 3 2 4 5",
	// 	"8 6 4 4 1",
	// 	"1 3 6 7 9",
	// }
	// input = sample

	safeCount := 0
	for _, line := range input {
		nums := SpaceSepToIntArr(line)
		if day02_safe_report(nums) {
			safeCount++
		}
	}

	fmt.Println("Day 02 Part 01:", safeCount)

	safeCount = 0
	for _, line := range input {
		nums := SpaceSepToIntArr(line)
		if day02_safe_with_removal(nums) {
			safeCount++
		}
	}

	fmt.Println("Day 02 Part 02:", safeCount)
}

func day02_safe_with_removal(nums []int) bool {
	if day02_safe_report(nums) {
		return true
	}

	for i := 0; i < len(nums); i++ {
		nums2 := make([]int, len(nums))
		copy(nums2, nums)
		nums2 = append(nums2[:i], nums2[i+1:]...)
		if day02_safe_report(nums2) {
			return true
		}
	}

	return false
}

func day02_safe_report(nums []int) bool {
	inc, dec, diff := true, true, true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= nums[i+1] {
			dec = false
		}
		if nums[i] >= nums[i+1] {
			inc = false
		}
		if Abs(nums[i]-nums[i+1]) < 1 || Abs(nums[i]-nums[i+1]) > 3 {
			diff = false
		}
	}
	return (inc || dec) && diff
}
