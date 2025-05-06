package main

import (
	"fmt"
	"slices"
)

// naive approach
func secondLargestElementNa(arr []int) int {
	temp := []int{} // create slices with distinct elements
	for _, v := range arr {
		if !slices.Contains(temp, v) {
			temp = append(temp, v)
		}
	}

	// sort arrays in descending order
	for i := range temp {
		for j := i + 1; j < len(temp); j++ {
			if temp[i] < temp[j] {
				temp[i], temp[j] = temp[j], temp[i]
			}
		}
	}

	if len(temp) == 1 {
		return -1
	}
	return temp[1]
}

func secondLargestElement(arr []int) int {
	l1 := 0
	l2 := -1
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[l1] {
			l2 = l1
			l1 = i
		} else if arr[i] != arr[l1] {
			if l2 == -1 || arr[i] > arr[l2] {
				l2 = i
			}
		}
	}
	if l2 == -1 {
		return -1
	}
	return arr[l2]
}

func main() {
	arr1 := []int{15, 5, 20}
	arr2 := []int{20, 10, 12, 5, 18, 20}
	arr3 := []int{20, 20}

	fmt.Println(secondLargestElementNa(arr1))
	fmt.Println(secondLargestElementNa(arr2))
	fmt.Println(secondLargestElementNa(arr3))

	fmt.Println(secondLargestElement(arr1))
	fmt.Println(secondLargestElement(arr2))
	fmt.Println(secondLargestElement(arr3))
}
