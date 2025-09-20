// 1. Two Sum
// Solved
// Easy
// Topics
// premium lock icon
// Companies
// Hint
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

// Constraints:

// 2 <= nums.length <= 104
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// Only one valid answer exists.

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	output := []int{0, 0}
	length := len(nums)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				output := []int{i, j}
				return output
			}
		}
	}
	return output
}

func main() {
	// Example 1:
	// Input: nums = [2,7,11,15], target = 9
	// Output: [1,2]
	input_1 := []int{2, 7, 11, 15}
	target_1 := 9
	fmt.Println(twoSum(input_1, target_1))
	// Example 2:
	// Input: nums = [3,2,4], target = 6
	// Output: [1,2]
	input_2 := []int{3, 2, 4}
	target_2 := 6
	fmt.Println(twoSum(input_2, target_2))
	// Example 3:
	// Input: nums = [3,3], target = 6
	// Output: [0,1]
	input_3 := []int{3, 3}
	target_3 := 6
	fmt.Println(twoSum(input_3, target_3))
}
