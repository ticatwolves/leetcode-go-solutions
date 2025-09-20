package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	given_int := x
	created_int := 0
	for x > 0 {
		reminder := x % 10
		created_int = (created_int * 10) + reminder
		x = x / 10
	}
	return created_int == given_int
}

func main() {
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(101))
	fmt.Println(isPalindrome(11010))
}
