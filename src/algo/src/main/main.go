package main

import "fmt"

func main() {
	fmt.Println("************************")

	A := []int{2, 5, 3, 2, 8, 1}
	B := []int{7, 9, 5, 2, 4, 10, 10}
	C := []int{6, 7, 5, 5, 3, 7}
	fmt.Printf("PresentInTwo: %d\n", PresentIntTwo(A, B, C))
	fmt.Println("************************")

	numsMax := []int{1, 2, 3, 4}
	fmt.Printf("Maximum Num %d\n", MaximumSubArray(numsMax))
	fmt.Println("************************")

	nums := []int{1, 2, 3, 4}
	fmt.Printf("ProductExcept: %d\n", ProductExceptSelf(nums))
	fmt.Println("************************")

	maxProductSub := []int{-5, 100, -3, -7}
	fmt.Printf("MaximumProductSubArray: %d\n", MaximumProductSubArray(maxProductSub))
	fmt.Println("************************")

	v1 := []int{7, 4, 1, 10}
	v2 := []int{4, 5, 8, 7}
	target := 17
	fmt.Printf("ClosestPair: %d\n", ClosestSumPair(v1, v2, target))
	fmt.Println("************************")

}
