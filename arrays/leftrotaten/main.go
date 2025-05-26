package main

import (
	"fmt"
	"slices"
)

func leftShiftD(arr []int, d int) []int {
	if d > len(arr) {
		d %= len(arr)
	}
	arr = slices.Concat(arr[d:], arr[:d])
	return arr
}
func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{10, 5, 30, 15}

	fmt.Println(leftShiftD(arr1, 2))
	fmt.Println(leftShiftD(arr2, 7))
}
