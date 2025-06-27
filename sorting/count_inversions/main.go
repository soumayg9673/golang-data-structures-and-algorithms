package main

import "fmt"

/*
	A pair (arr[i], arr[j])
	forms an inversion when
		i<j
	and arr[i] > arr[j]
*/

func countInversionNa(arr []int) int {
	total := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				total++
			}
		}
	}
	return total
}

func main() {
	arr1 := []int{2, 4, 1, 3, 5}  // 3
	arr2 := []int{10, 20, 30, 40} // 0
	arr3 := []int{40, 30, 20, 10} // 6

	// Naive approach
	fmt.Println(countInversionNa(arr1))
	fmt.Println(countInversionNa(arr2))
	fmt.Println(countInversionNa(arr3))

}
