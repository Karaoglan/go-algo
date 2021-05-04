package main

import (
	"fmt"
)

func main() {
	A := []int{2, 5, 3, 2, 8, 1}
	B := []int{7, 9, 5, 2, 4, 10, 10}
	C := []int{6, 7, 5, 5, 3, 7}
	fmt.Println(PresentIntTwo(A, B, C))
}

func PresentIntTwo(A, B, C []int) []int {

	intMap := make(map[int]int)

	findIntersectAndPutMap(A, B, intMap)
	findIntersectAndPutMap(B, C, intMap)
	findIntersectAndPutMap(A, C, intMap)

	keys := make([]int, len(intMap))
	count := 0
	for k := range intMap {
		keys[count] = k
		count++
	}
	fmt.Println(count)

	return keys
}

func putIntoMap(intersectionMap map[int]int, value int) {
	count := 0

	if val, ok := intersectionMap[value]; ok {
		count = val
	}
	count++
	intersectionMap[value] = count
}

func findIntersectAndPutMap(A, B []int, intersectionMap map[int]int) {
	hash := make(map[int]bool)

	for i := 0; i < len(A); i++ {
		hash[A[i]] = true
	}

	for i := 0; i < len(B); i++ {
		if _, found := hash[B[i]]; found {
			putIntoMap(intersectionMap, B[i])
		}
	}
}