package main

import "fmt"

func union(arr, brr []int) {
	trr := []int{}
	an, bn := len(arr), len(brr)
	i, j := 0, 0
	for i < an && j < bn {
		if i > 0 && arr[i] == arr[i-1] {
			i++
			continue
		}
		if j > 0 && brr[j] == brr[j-1] {
			j++
			continue
		}
		if arr[i] > brr[j] {
			trr = append(trr, brr[j])
			j++
		} else if arr[i] < brr[j] {
			trr = append(trr, arr[i])
			i++
		} else {
			trr = append(trr, arr[i])
			i++
			j++
		}
	}
	for ; i < an; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		trr = append(trr, arr[i])
	}
	for ; j < bn; j++ {
		if j > 0 && brr[j] == brr[j-1] {
			continue
		}
		trr = append(trr, brr[j])
	}
	fmt.Println(trr)
}

func main() {
	a1 := []int{3, 5, 8}
	b1 := []int{2, 8, 9, 10, 15}
	union(a1, b1)

	a2 := []int{2, 3, 3, 3, 4, 4}
	b2 := []int{4, 4}
	union(a2, b2)

}
