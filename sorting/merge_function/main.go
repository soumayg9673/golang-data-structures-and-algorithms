package main

import (
	"fmt"
)

func sortFuncNa(arr []int, l, m, h int) {
	for i := m + 1; i <= h; i++ {
		for arr[m] > arr[i] {
			j := l
			for ; j <= m; j++ {
				if arr[j] > arr[i] {
					break
				}

			}
			arr[j], arr[i] = arr[i], arr[j]
		}
		m++
	}
	fmt.Println(arr)
}

func sortFunc(arr []int, l, m, h int) {

	ln, rn := m-l+1, h-m

	var lrr, rrr []int

	for t := range ln {
		lrr = append(lrr, arr[t])
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
	fmt.Println(arr)
}

func main() {
	sortFuncNa([]int{10, 15, 20, 11, 30}, 0, 2, 4)
	sortFuncNa([]int{5, 8, 12, 14, 7}, 0, 3, 4)
	sortFuncNa([]int{10, 15, 20, 40, 8, 11, 55}, 0, 3, 6)

	sortFunc([]int{10, 15, 20, 11, 30}, 0, 2, 4)
	sortFunc([]int{5, 8, 12, 14, 7}, 0, 3, 4)
	sortFunc([]int{10, 15, 20, 40, 8, 11, 55}, 0, 3, 6)
}
