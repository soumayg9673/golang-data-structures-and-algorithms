package main

import "fmt"

func naivePartition(arr []int, l, h, p int) {
	tarr := make([]int, len(arr))
	idx := 0
	// add elements less than or equal to p-th element
	for i := l; i <= h; i++ {
		if arr[i] <= arr[p] && i != p {
			tarr[idx] = arr[i]
			idx++
		}
	}

	// add p-th element to temporary array
	tarr[idx] = arr[p]
	idx++

	// add elements greater than p-th element
	for i := l; i <= h; i++ {
		if arr[i] > arr[p] {
			tarr[idx] = arr[i]
			idx++
		}
	}
	fmt.Println(tarr)
}

func main() {
	arr1 := []int{3, 8, 6, 12, 10, 7}
	arr2 := []int{5, 13, 6, 9, 12, 11, 8}

	naivePartition(arr1, 0, len(arr1)-1, 5)
	naivePartition(arr2, 0, len(arr2)-1, 6)
}
