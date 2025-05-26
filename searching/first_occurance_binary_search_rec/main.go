package main

import "fmt"

func firstOcc(arr []int, x, l, h int) int {
	if l > h {
		return -1
	}
	m := (l + h) / 2
	if arr[m] > x {
		return firstOcc(arr, x, l, m-1)
	} else if arr[m] < x {
		return firstOcc(arr, x, m+1, h)
	} else {
		if m == 0 || arr[m-1] != arr[m] {
			return m
		} else {
			return firstOcc(arr, x, l, m-1)
		}
	}
}

func main() {
	arr1 := []int{1, 10, 10, 10, 20, 20, 40}
	arr2 := []int{10, 20, 30}
	arr3 := []int{15, 15, 15}

	fmt.Println(firstOcc(arr1, 10, 0, len(arr1)-1))
	fmt.Println(firstOcc(arr2, 15, 0, len(arr2)-1))
	fmt.Println(firstOcc(arr3, 15, 0, len(arr3)-1))
}
