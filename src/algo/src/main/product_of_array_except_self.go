package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(ProductExceptSelf(nums))
}

func ProductExceptSelf(nums []int) []int {
	sol := make([]int, len(nums))

	if len(nums) == 0 {
		return sol
	}

	sol[0] = 1
	for i := 1; i < len(nums); i++ {
		sol[i] = sol[i-1] * nums[i-1]
	}
	rightProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		sol[i] = sol[i] * rightProduct
		rightProduct *= nums[i]
	}
	return sol
}