package main

func FindMinInSorted(nums []int) int {
	lo := 0
	hi := len(nums) - 1
	for lo < hi {
		guess := (lo + hi) / 2
		if nums[guess] > nums[hi] {
			lo = guess + 1
		} else {
			hi = guess
		}
	}
	return nums[lo]
}
