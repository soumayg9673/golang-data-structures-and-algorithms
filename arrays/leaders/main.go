package main

import (
	"fmt"
	"slices"
)

func leaders(arr []int) []int {
	tar := []int{}

	// insert last element
	tar = append(tar, arr[len(arr)-1])

	for i := len(arr) - 2; i >= 0; i-- {
		if tar[len(tar)-1] < arr[i] {
			tar = append(tar, arr[i])
		}
	}

	slices.Reverse(tar)
	return tar
}

func main() {
	arr1 := []int{7, 10, 4, 3, 6, 5, 2}
	arr2 := []int{7, 10, 4, 10, 6, 5, 2}
	arr3 := []int{10, 20, 30}
	arr4 := []int{30, 20, 10}

	fmt.Println(leaders(arr1))
	fmt.Println(leaders(arr2))
	fmt.Println(leaders(arr3))
	fmt.Println(leaders(arr4))
}
