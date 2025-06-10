package main

import "fmt"

func merge(arr []int, l, m, h int) {

	ln, rn := m-l+1, h-m

	var lrr, rrr []int

	for t := range ln {
		lrr = append(lrr, arr[l+t])
	}
	for t := range rn {
		rrr = append(rrr, arr[t+m+1])
	}

	i, j, k := 0, 0, l
	for i < ln && j < rn {
		if lrr[i] <= rrr[j] {
			arr[k] = lrr[i]
			i++
			k++
		} else {
			arr[k] = rrr[j]
			j++
			k++
		}

	}
	for ; j < rn; j++ {
		arr[k] = rrr[j]
		k++
	}
	for ; i < ln; i++ {
		arr[k] = lrr[i]
		k++
	}
}

func mergeSort(arr []int, l, r int) {
	if r > l {
		m := l + ((r - l) / 2)
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)
		merge(arr, l, m, r)
	}
}
func main() {
	arr := []int{30, 10, 18, 3, 2, 16, 50}
	mergeSort(arr, 0, 6)
	fmt.Println(arr)
}
