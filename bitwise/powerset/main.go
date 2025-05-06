package main

import (
	"fmt"
	"strings"
)

func powerSetNa(s string) string {
	arr := []string{}
	arr = append(arr, "")
	c := 1 << len(s)
	for i := 1; i < c; i++ {
		temp := []string{}
		counter := 1
		for j := i; j > 0; counter++ {
			if j%2 == 1 {
				temp = append(temp, string(s[counter-1]))
			}
			j /= 2
		}
		arr = append(arr, strings.Join(temp, ""))
	}
	return strings.Join(arr, ",")
}

func powerSet(s string) string {
	ln := len(s)
	c := 1 << ln
	arr := []string{}
	for i := range c {
		temp := []string{}
		for j := range ln {
			if i&(1<<j) != 0 {
				temp = append(temp, string(s[j]))
			}
		}
		arr = append(arr, strings.Join(temp, ""))
	}
	return strings.Join(arr, ",")
}

func main() {
	// naive approach
	fmt.Println(powerSetNa("abc"))

	fmt.Println(powerSet("abc"))
}
