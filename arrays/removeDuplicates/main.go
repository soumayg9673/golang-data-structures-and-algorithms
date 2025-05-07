package main

import (
	"fmt"
	"slices"
)

func removeDuplicates(arr []int) []int {
	t := []int{}
	for _, v := range arr {
		if !slices.Contains(t, v) {
			t = append(t, v)
		}
	}
	return t
}

func main() {
	arr1 := []int{15}
	arr2 := []int{10, 20, 20, 30, 30, 30, 30, 30}
	arr3 := []int{20, 20}

	fmt.Println(removeDuplicates(arr1))
	fmt.Println(removeDuplicates(arr2))
	fmt.Println(removeDuplicates(arr3))
}
