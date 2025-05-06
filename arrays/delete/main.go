package main

import (
	"fmt"
	"slices"
)

func deleteElement(arr []int, x int) []int {
	idx := slices.Index(arr, x)
	if idx == -1 {
		return arr
	}
	arr = slices.Concat(arr[0:idx], arr[idx+1:])
	return arr
}

func deleteIndex(arr []int, pos int) []int {
	arr = slices.Concat(arr[0:pos-1], arr[pos:])
	return arr
}

func main() {
	arr1 := []int{5, 10, 15, 20}
	fmt.Println(deleteElement(arr1, 10))
	fmt.Println(deleteElement(arr1, 20))

	arr2 := []int{1, 2, 3, 4}
	fmt.Println(deleteIndex(arr2, 2))
	fmt.Println(deleteIndex(arr2, 4))
}
