package main

import (
	"fmt"
)

func maximumDifferenceNa(arr []int) int {
	md := arr[1] - arr[0]
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			temp := arr[j] - arr[i]
			if temp > md {
				md = temp
			}
		}
	}
	return md
}

func maximumDifference(arr []int) int {
	res := arr[1] - arr[0]
	min := arr[0]

	for j := 1; j < len(arr); j++ {
		diff := arr[j] - min
		if diff > res {
			res = diff
		}
		if arr[j] <= min {
			min = arr[j]
		}
	}
	return res
}

func main() {
	arr1 := []int{2, 3, 10, 6, 4, 8, 1}
	arr2 := []int{7, 9, 5, 6, 3, 2}
	arr3 := []int{10, 20, 30}
	arr4 := []int{30, 10, 8, 2}

	fmt.Println("Naive Approach")
	fmt.Println(maximumDifferenceNa(arr1))
	fmt.Println(maximumDifferenceNa(arr2))
	fmt.Println(maximumDifferenceNa(arr3))
	fmt.Println(maximumDifferenceNa(arr4))

	fmt.Println("Optimized Approach")
	fmt.Println(maximumDifference(arr1))
	fmt.Println(maximumDifference(arr2))
	fmt.Println(maximumDifference(arr3))
	fmt.Println(maximumDifference(arr4))
}
