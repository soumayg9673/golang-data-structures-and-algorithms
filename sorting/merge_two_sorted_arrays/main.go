package main

import (
	"fmt"
	"slices"
)

/*
	divide and conquer algorithm
	stable algorithm
	Theta(n log n) time and O(n) aux space
	well suited for linked lists. works in O(1) aux space
	used in external sorting
	in general for arrays, QUICK SORT outperforms
*/

func mergeSortNa(arr []int, brr []int) {
	trr := slices.Concat(arr, brr)
	slices.Sort(trr)
	fmt.Println(trr)
}

func mergeSort(arr []int, brr []int) {
	trr := []int{}
	i, j := 0, 0
	for i < len(arr) && j < len(brr) {
		if arr[i] <= brr[j] {
			trr = append(trr, arr[i])
			i++
		} else {
			trr = append(trr, brr[j])
			j++
		}
	}
	for ; j < len(brr); j++ {
		trr = append(trr, brr[j])
	}
	for ; i < len(arr); i++ {
		trr = append(trr, arr[i])
	}
	fmt.Println(trr)
}

func main() {
	a1 := []int{10, 15, 20, 20}
	b1 := []int{1, 2}
	mergeSortNa(a1, b1)
	mergeSort(a1, b1)

	a2 := []int{10, 20, 35}
	b2 := []int{5, 50, 50}
	mergeSortNa(a2, b2)
	mergeSort(a2, b2)
}
