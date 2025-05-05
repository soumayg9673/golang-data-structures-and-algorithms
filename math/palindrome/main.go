package main

import "fmt"

func palindrome(n int) bool {
	tempN := n
	reverseN := 0
	for tempN > 0 {
		reverseN = (reverseN * 10) + tempN%10
		tempN /= 10
	}
	return reverseN == n
}

func main() {
	fmt.Println(palindrome(766))
	fmt.Println(palindrome(9673769))
}
