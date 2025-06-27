package main

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

func main() {
	arr := []int{5, 3, 8, 4, 2, 7, 1, 10}

	hoarePartition(arr, 0, len(arr)-1)
}
