package main

import (
	"fmt"
	"slices"
)

// ascending order
func checkAscSort(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

// descending order
func checkDescSort(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			return false
		}
	}
	return true
}

func main() {
	arr1 := []int{15}
	arr2 := []int{5, 18, 20}
	arr3 := []int{20, 20}
	arr4 := []int{20, 10}

	fmt.Println(checkAscSort(arr1))
	fmt.Println(checkAscSort(arr2))
	fmt.Println(checkAscSort(arr3))
	fmt.Println(checkAscSort(arr4))

	fmt.Println(checkDescSort(arr1))
	fmt.Println(checkDescSort(arr2))
	fmt.Println(checkDescSort(arr3))
	fmt.Println(checkDescSort(arr4))

	// use in-built function
	fmt.Println(slices.IsSorted(arr1))
	fmt.Println(slices.IsSorted(arr2))
	fmt.Println(slices.IsSorted(arr3))
	fmt.Println(slices.IsSorted(arr4))
}
