package main

import "fmt"

func binarySearch(arr []int, x int) int {
	n := len(arr)
	l, h := 0, n-1
	for l <= h {
		m := (l + h) / 2
		if arr[m] == x {
			return m
		} else if arr[m] > x {
			h = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

func main() {
	arr1 := []int{10, 20, 30, 40, 50, 60, 70}
	arr2 := []int{5, 15, 25}
	arr3 := []int{5, 10, 15, 20}
	arr4 := []int{10, 10}

	fmt.Println(binarySearch(arr1, 40))
	fmt.Println(binarySearch(arr1, 20))
	fmt.Println(binarySearch(arr2, 25))
	fmt.Println(binarySearch(arr3, 5))
	fmt.Println(binarySearch(arr4, 10))
	fmt.Println(binarySearch(arr4, 1))
}
