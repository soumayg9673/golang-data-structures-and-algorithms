package main

import "fmt"

// This can work for any type array (sorted/unsorted)
func intersectionNa(arr, brr []int) {
	trr := []int{}
	aln, bln := len(arr), len(brr)
	for i := range aln {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		for j := range bln {
			if arr[i] == brr[j] {
				trr = append(trr, arr[i])
				break
			}
		}
	}
	fmt.Println(trr)
}

func intersection(arr, brr []int) {
	trr := []int{}
	for i, j := 0, 0; i < len(arr) && j < len(brr); {
		if i > 0 && arr[i] == arr[i-1] {
			i++
			continue
		}
		if arr[i] > brr[j] {
			j++
		} else if arr[i] < brr[j] {
			i++
		} else {
			trr = append(trr, arr[i])
			i++
			j++
		}
	}
	fmt.Println(trr)
}

func main() {
	a1 := []int{1, 20, 20, 40, 60}
	b1 := []int{2, 20, 20, 20}
	intersectionNa(a1, b1)
	intersection(a1, b1)
}
