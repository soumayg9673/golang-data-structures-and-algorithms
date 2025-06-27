package main

import "fmt"

// Considering pivot the last element of the array
func lomutoParition(arr []int, l, h int) int {
	p := arr[h]
	i := l - 1
	for j := l; j <= h-1; j++ {
		if arr[j] < p {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[h] = arr[h], arr[i+1]
	return i + 1
}

func quickSort(arr []int, l, h int) {
	if l < h {
		p := lomutoParition(arr, l, h)
		quickSort(arr, l, p-1)
		quickSort(arr, p+1, h)
	}
}

func main() {
	arr1 := []int{10, 80, 30, 90, 40, 50, 70}

	quickSort(arr1, 0, len(arr1)-1)
	fmt.Println(arr1)
}
