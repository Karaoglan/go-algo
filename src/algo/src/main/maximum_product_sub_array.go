package main

import (
	"github.com/pkg/math"
)

func MaximumProductSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	temp := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		temp = math.Max(temp*nums[i], nums[i])
		sum = math.Max(sum, temp)
	}

	return sum
}

func MaximumProductSubBruteForce(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
	}
	// TODO
	return 0
}
