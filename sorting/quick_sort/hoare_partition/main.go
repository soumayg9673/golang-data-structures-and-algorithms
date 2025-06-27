package main

import "fmt"

// Considering the pivot element as the first element of the array
func hoarePartition(arr []int, l, h int) int {
	p := arr[l]
	i, j := l-1, h+1
	for {
		for i++; arr[i] < p; i++ {
		}
		for j--; arr[j] > p; j-- {
		}
		if i >= j {
			return j
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func quickSort(arr []int, l, h int) {
	if l < h {
		p := hoarePartition(arr, l, h)
		quickSort(arr, l, p)
		quickSort(arr, p+1, h)
	}
}

func main() {
	arr := []int{5, 3, 8, 4, 2, 7, 1, 10}

	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
