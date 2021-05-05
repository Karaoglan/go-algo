package main

import (
	"github.com/pkg/math"
)

func MaximumProductSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// for positive values
	maxEndingHere := 1

	// for negative values
	minEndingHere := 1

	maxSoFar := 0
	flag := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			maxEndingHere *= nums[i]
			minEndingHere = math.Min(minEndingHere*nums[i], 1)
			flag = 1

		} else if nums[i] == 0 {
			maxEndingHere = 1
			minEndingHere = 1

		} else {
			temp := maxEndingHere
			maxEndingHere = math.Max(minEndingHere*nums[i], 1)
			minEndingHere = temp * nums[i]
		}

		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
	}

	if flag == 0 && maxSoFar == 0 {
		return 0
	}
	return maxSoFar
}

func MaximumProductSubBruteForce(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]

	for i := 0; i < len(nums); i++ {
		mul := nums[i]
		for j := i + 1; j < len(nums); j++ {
			mul *= nums[j]
		}
		result = math.Max(result, mul)
	}
	return result
}
