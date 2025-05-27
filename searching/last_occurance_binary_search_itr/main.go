package main

import "fmt"

func lastOcc(arr []int, x int) int {
	n := len(arr)
	l, h := 0, len(arr)-1
	for l <= h {
		m := (l + h) / 2
		if arr[m] > x {
			h = m - 1
		} else if arr[m] < x {
			l = m + 1
		} else {
			if m == n-1 || arr[m+1] != arr[m] {
				return m
			} else {
				l = m + 1
			}
		}
	}
	return -1
}
func main() {
	arr1 := []int{10, 15, 20, 20, 40, 40}
	arr2 := []int{5, 8, 8, 10, 10}
	arr3 := []int{8, 10, 10, 12}

	fmt.Println(lastOcc(arr1, 20))
	fmt.Println(lastOcc(arr2, 10))
	fmt.Println(lastOcc(arr3, 7))
}
