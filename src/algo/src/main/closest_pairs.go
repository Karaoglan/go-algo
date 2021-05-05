package main

import (
	"fmt"
	"sort"
)

func ClosestSumPair(a1, a2 []int, target int) []int {
	a1Sorted := make([]int, len(a1))
	copy(a1Sorted, a1)
	sort.Ints(a1Sorted)
	fmt.Printf("a1Sorted: %d\n", a1Sorted)

	a2Sorted := make([]int, len(a2))
	copy(a2Sorted, a2)
	sort.Ints(a2Sorted)
	fmt.Printf("a2Sorted: %d\n", a2Sorted)

	i := 0
	j := len(a2Sorted) - 1
	smallestDiff := abs(a1Sorted[0] + a2Sorted[0] - target)
	closestPair := []int{a1Sorted[0], a2Sorted[0]}
	for i < len(a1Sorted) && j >= 0 {
		fmt.Printf("i : %d\n", i)
		fmt.Printf("j : %d\n", j)

		v1 := a1Sorted[i]
		v2 := a2Sorted[j]
		currentDiff := v1 + v2 - target
		fmt.Printf("currentDif : %d\n", currentDiff)
		fmt.Printf("smallestDif : %d\n", smallestDiff)
		if abs(currentDiff) < smallestDiff {
			smallestDiff = abs(currentDiff)
			closestPair[0] = v1
			closestPair[1] = v2
		}
		if currentDiff == 0 {
			return closestPair
		} else if currentDiff < 0 {
			i += 1
		} else {
			j -= 1
		}

		fmt.Println("-----")

	}
	return closestPair
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
