package main

import "fmt"

func binarySearch(arr []int, x, l, h int) int {
	if l > h {
		return -1
	}
	m := (l + h) / 2
	if arr[m] == x {
		return m
	}
	if arr[m] > x {
		return binarySearch(arr, x, l, m-1)
	} else {
		return binarySearch(arr, x, m+1, h)
	}
}

func main() {
	arr1 := []int{10, 20, 30, 40, 50}
	arr2 := []int{10, 20, 30}
	arr3 := []int{10, 20, 30, 40}
	arr4 := []int{10, 20}
	arr5 := []int{15}

	fmt.Println(binarySearch(arr1, 20, 0, len(arr1)-1))
	fmt.Println(binarySearch(arr2, 30, 0, len(arr2)-1))
	fmt.Println(binarySearch(arr2, 10, 0, len(arr2)-1))
	fmt.Println(binarySearch(arr3, 15, 0, len(arr3)-1))
	fmt.Println(binarySearch(arr4, 5, 0, len(arr4)-1))
	fmt.Println(binarySearch(arr5, 20, 0, len(arr5)-1))
}
