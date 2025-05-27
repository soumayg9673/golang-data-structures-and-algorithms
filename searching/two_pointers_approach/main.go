package main

import "fmt"

/*
	Only for sorted array (ascending order)
*/

func checkPair(arr []int, x int) bool {
	l, h := 0, len(arr)-1
	for l != h {
		if arr[l]+arr[h] == x {
			return true
		} else if arr[l]+arr[h] > x {
			h--
		} else {
			l++
		}
	}
	return false
}
func main() {
	arr1 := []int{2, 4, 8, 9, 11, 12, 20, 30}

	fmt.Println(checkPair(arr1, 23))
}
