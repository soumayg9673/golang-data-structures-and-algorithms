package main

import "fmt"

func firstOcc(arr []int, x int) int {
	l, h := 0, len(arr)-1
	for l <= h {
		m := (l + h) / 2
		if arr[m] > x {
			h = m - 1
		} else if arr[m] < x {
			l = m + 1
		} else {
			if m == 0 || arr[m-1] != arr[m] {
				return m
			} else {
				h = m - 1
			}
		}
	}
	return -1
}
func main() {
	arr1 := []int{1, 10, 10, 10, 20, 20, 40}
	arr2 := []int{10, 20, 30}
	arr3 := []int{15, 15, 15}

	fmt.Println(firstOcc(arr1, 10))
	fmt.Println(firstOcc(arr2, 15))
	fmt.Println(firstOcc(arr3, 15))
}
