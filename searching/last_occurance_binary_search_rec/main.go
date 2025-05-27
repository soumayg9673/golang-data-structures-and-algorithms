package main

import "fmt"

func lastOcc(arr []int, x, l, h int) int {
	if l > h {
		return -1
	}
	m := (l + h) / 2
	if arr[m] > x {
		return lastOcc(arr, x, l, m-1)
	} else if arr[m] < x {
		return lastOcc(arr, x, m+1, h)
	} else {
		if m == len(arr)-1 || arr[m+1] != arr[m] {
			return m
		} else {
			return lastOcc(arr, x, m+1, h)
		}
	}
}

func main() {
	arr1 := []int{10, 15, 20, 20, 40, 40}
	arr2 := []int{5, 8, 8, 10, 10}
	arr3 := []int{8, 10, 10, 12}

	fmt.Println(lastOcc(arr1, 20, 0, len(arr1)-1))
	fmt.Println(lastOcc(arr2, 10, 0, len(arr2)-1))
	fmt.Println(lastOcc(arr3, 7, 0, len(arr3)-1))
}
