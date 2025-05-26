package main

import (
	"fmt"
	"math"
	"slices"
)

func calculate(arr []int) int {
	// check the array for 1
	i1 := slices.Index(arr, 1)
	if i1 == -1 {
		return 0
	}
	temp := 1
	res := 1
	for i := i1 + 1; i < len(arr); i++ {
		switch arr[i] {
		case 0:
			temp = 0
		case 1:
			temp++
			res = int(math.Max(float64(temp), float64(res)))
		}
	}
	return res
}

func main() {
	arr1 := []int{0, 1, 1, 0, 1, 0}
	arr2 := []int{1, 1, 1, 1}
	arr3 := []int{0, 0, 0}
	arr4 := []int{1, 0, 1, 1, 1, 1, 0, 1, 1}

	fmt.Println(calculate(arr1))
	fmt.Println(calculate(arr2))
	fmt.Println(calculate(arr3))
	fmt.Println(calculate(arr4))
}
