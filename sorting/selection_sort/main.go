package main

import (
	"fmt"
	"math"
)

/*
	Not stable algorithm
*/

func selectionSortNa(arr []int) {
	tarr := []int{}
	for i := range len(arr) {
		mi := 0
		for j := 1; j < len(arr); j++ {
			if arr[j] < arr[mi] {
				mi = j
			}
		}
		tarr = append(tarr, arr[mi])
		tarr[i] = arr[mi]
		arr[mi] = math.MaxInt
	}
	fmt.Println(tarr)
}

func selectionSortMemOpt(arr []int) {
	for i := range len(arr) - 1 {
		mi := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[mi] {
				mi = j
			}
		}
		arr[i], arr[mi] = arr[mi], arr[i]
	}
	fmt.Println(arr)
}

func main() {
	selectionSortNa([]int{10, 5, 8, 20, 2, 18})
	selectionSortMemOpt([]int{10, 5, 8, 20, 2, 18})
}
