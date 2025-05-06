package main

import "fmt"

func palindrome(str string, s int, e int) bool {
	if s >= e {
		return true
	}
	return str[s] == str[e] && palindrome(str, s+1, e-1)
}

func main() {
	fmt.Println(palindrome("aba", 0, len("aba")-1))
	fmt.Println(palindrome("a", 0, len("a")-1))
	fmt.Println(palindrome("ab", 0, len("ab")-1))
}
