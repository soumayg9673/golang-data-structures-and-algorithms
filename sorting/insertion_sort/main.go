package main

import "fmt"

/*
	O(n2) worst case
	In-Place and Stable
	Used in practice for small arrays
	O(n) best case
*/

func insertionSortNa(arr []int) {
	for i := 1; i < len(arr); i++ {
		k := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > k; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = k
	}
	fmt.Println(arr)
}

func main() {
	insertionSortNa([]int{20, 5, 40, 60, 10, 30})
}
