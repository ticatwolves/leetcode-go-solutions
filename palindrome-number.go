// 9. Palindrome Number
// Solved
// Easy
// Topics
// premium lock icon
// Companies
// Hint
// Given an integer x, return true if x is a palindrome, and false otherwise.

// Example 1:

// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
// Example 2:

// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
// Example 3:

// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

// Constraints:

// -231 <= x <= 231 - 1

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
